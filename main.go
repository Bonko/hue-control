package main

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go assets/

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gbbr/hue"
)

const (
	port = 9091
)

func main() {
	log.Printf("listening on port %d", port)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/sleepTimer", startSleepTimer)
	port_string := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(port_string, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("reached end")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	debugHtmlPath, _ := filepath.Abs(".debug/html/index.html")

	if _, err := os.Stat(debugHtmlPath); err == nil {
		t := template.Must(template.ParseFiles(debugHtmlPath))
		t.Execute(w, nil)
	} else {
		html, err := Asset("assets/index.html")
		if err != nil {
			http.Error(w, "Could not load html: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(html)
	}
}

func auth() *hue.Bridge {
	b, err := hue.Discover()
	if err != nil {
		log.Println("discovery failed")
		log.Fatal(err)
	}
	if !b.IsPaired() {
		log.Println("program not paired with bridge, pairing now")
		// link button must be pressed before calling
		if err := b.Pair(); err != nil {
			log.Println("pairing failed")
			log.Fatal(err)
		}
	}
	return b
}

func startSleepTimer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	d := r.Form["duration"][0]
	duration, err := strconv.Atoi(d)
	if err != nil {
		log.Fatal("cannot convert duration to int: %s", err)
	}

	b := r.Form["brightness"][0]
	b64, err := strconv.ParseUint(b, 10, 8)
	if err != nil {
		log.Fatal("cannot convert brightness to uint: %s", err)
	}
	brightness := uint8(b64)
	go sleepTimer(time.Duration(duration), brightness)
	fmt.Fprintf(w, "Started sleepTimer (duration: %d minutes)", duration)
}

func calcSteps(start uint8) uint8 {
	return start / 10
}

func getAvailableLights(b *hue.Bridge) ([]*hue.Light, error) {
	lightsService := b.Lights()
	availableLights, err := lightsService.List()
	if err != nil {
		return nil, err
	}
	return availableLights, nil
}

func getAvailableLightNames(lights []*hue.Light) []string {
	var availableLightNames []string
	for _, v := range lights {
		availableLightNames = append(availableLightNames, v.Name)
	}
	return availableLightNames
}

func sleepTimer(duration time.Duration, startBrightness uint8) {
	b := auth()
	availableLights, err := getAvailableLights(b)
	if err != nil {
		log.Fatalf("error while retrieving lights: %s", err)
	}
	availableLightNames := getAvailableLightNames(availableLights)
	log.Printf("availableLightNames: %v", availableLightNames)
	return
	nk, err := b.Lights().Get("Nachtkaestchen")
	if err != nil {
		log.Fatal(err)
	}
	original_brightness := nk.State.Brightness
	log.Println("Current brightness:", original_brightness)

	log.Println("Turning light on")
	nk.On()

	var brightness uint8

	if startBrightness == 0 {
		brightness = original_brightness
		log.Println("Keeping brightness:", brightness)
	} else {
		brightness = startBrightness
		log.Println("Setting brightness to:", brightness)
		nk.Set(&hue.State{
			Brightness: brightness,
		})
	}

	brightnessDecreaseStep := calcSteps(brightness)
	remaining_time := duration * time.Minute
	interval := remaining_time / 10

	for remaining_time > 0 && brightness > brightnessDecreaseStep {
		log.Println("Sleeping", interval)
		time.Sleep(interval)
		brightness = brightness - brightnessDecreaseStep
		nk.Set(&hue.State{
			Brightness: brightness,
		})
		log.Println("Decreased brightness to:", nk.State.Brightness, "expected:", brightness)
		remaining_time = remaining_time - interval
		log.Println("Remaining Time:", remaining_time)
	}

	log.Println("Setting brightness back to original value:", original_brightness)
	// use a high transition time (10 seconds) in order to set brightness back to its original value
	// without actually raising the brightness (and wake me up again)
	// it has to be done this way, because brightness cannot be set when the light is turned off
	nk.Set(&hue.State{
		Brightness:     original_brightness,
		TransitionTime: 100,
	})

	log.Println("Turning light off")
	nk.Off()
}
