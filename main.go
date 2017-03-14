package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gbbr/hue"
)

func main() {
	log.Println("listening..")
	//	sleepTimer()
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/sleepTimer", startSleepTimer)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("reached end")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	html := `<a href="/sleepTimer">sleep timer</a>`
	fmt.Fprintf(w, html)
}

func auth() *hue.Bridge {
	b, err := hue.Discover()
	if err != nil {
		log.Println("discovery failed")
		log.Fatal(err)
	}
	if !b.IsPaired() {
		log.Println("program not paired with bridge")
		// link button must be pressed before calling
		if err := b.Pair(); err != nil {
			log.Println("pairing failed")
			log.Fatal(err)
		}
	}
	return b
}

func startSleepTimer(w http.ResponseWriter, r *http.Request) {
	go sleepTimer()
	fmt.Fprintf(w, "Started sleepTimer")
}
func sleepTimer() {
	b := auth()
	nk, err := b.Lights().Get("Nachtkaestchen")
	if err != nil {
		log.Fatal(err)
	}
	brightness := uint8(255)
	log.Println("Turning light on")
	nk.On()
	nk.Set(&hue.State{
		Brightness: brightness,
	})
	remaining_time := 20 * time.Minute
	for remaining_time > 0 && brightness > 25 {
		interval := 2 * time.Minute
		log.Println("Sleeping", interval)
		time.Sleep(interval)
		brightness = brightness - 25
		log.Println("Decreasing brightness to", brightness)
		nk.Set(&hue.State{
			Brightness: brightness,
		})
		log.Println(nk.State.Brightness)
		remaining_time = remaining_time - interval
		log.Println("Remaining Time:", remaining_time)
	}
	nk.Off()
}
