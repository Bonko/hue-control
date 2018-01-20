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

	"github.com/Bonko/hue-control/controller"
	"github.com/gbbr/hue"
)

const (
	port = 9091
)

var (
	// FIXME: inject ctrl to handlers instead of using global var
	ctrl *controller.Controller
)

func init() {
	// FIXME: how to catch error with getting an
	// "ctrl declared and not used" error?
	ctrl, _ = controller.NewController()
}

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

	availableLights, err := ctrl.AvailableLights()
	if err != nil {
		log.Fatalf("error while retrieving lights: %s", err)
	}

	if _, err := os.Stat(debugHtmlPath); err == nil {
		t := template.Must(template.ParseFiles(debugHtmlPath))
		t.Execute(w, availableLights)
	} else {
		html, err := Asset("assets/index.html")
		if err != nil {
			http.Error(w, "Could not load html: "+err.Error(), http.StatusInternalServerError)
			return
		}
		// transform bytes from asset into string
		htmlAsString := string(html[:])
		t := template.Must(template.New("name").Parse(htmlAsString))
		t.Execute(w, availableLights)
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
	lightName := r.Form["light"][0]
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
	go sleepTimer(lightName, time.Duration(duration), brightness)
	fmt.Fprintf(w, "Started sleepTimer (duration: %d minutes)", duration)
}

func calcSteps(start uint8) uint8 {
	return start / 10
}

func sleepTimer(lightName string, duration time.Duration, startBrightness uint8) {
	original_brightness, err := ctrl.Brightness(lightName)
	if err != nil {
		log.Fatalf("Error: could not get current brightness: %s", err)
	}
	log.Println("Current brightness:", original_brightness)

	log.Println("Turning light on")
	if err := ctrl.On(lightName); err != nil {
		log.Fatalf("Error: could not turn light on: %s", err)
	}

	var brightness uint8

	if startBrightness == 0 {
		brightness = original_brightness
		log.Println("Keeping brightness:", brightness)
	} else {
		brightness = startBrightness
		log.Println("Setting brightness to:", brightness)
		if err := ctrl.SetBrightness(lightName, brightness); err != nil {
			log.Fatal(err)
		}
	}

	brightnessDecreaseStep := calcSteps(brightness)
	remaining_time := duration * time.Minute
	interval := remaining_time / 10

	for remaining_time > 0 && brightness > brightnessDecreaseStep {
		log.Println("Sleeping", interval)
		time.Sleep(interval)
		realBrightness, err := ctrl.Brightness(lightName)
		if err != nil {
			log.Fatal(err)
		}
		if realBrightness != brightness {
			log.Println("Light was changed externally, cancelling timer")
			return
		}

		brightness = brightness - brightnessDecreaseStep
		if err := ctrl.SetBrightness(lightName, brightness); err != nil {
			log.Fatal(err)
		}

		log.Println("Decreased brightness to:", brightness)
		remaining_time = remaining_time - interval
		log.Println("Remaining Time:", remaining_time)
	}

	log.Printf("Setting brightness back to original value(%d) and turning off light", original_brightness)
	if err := ctrl.TurnOffAndRestoreBrightness(lightName, original_brightness); err != nil {
		log.Fatal(err)
	}
}
