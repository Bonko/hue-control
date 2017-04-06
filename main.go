package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Bonko/hue"
)

func main() {
	log.Println("listening..")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/sleepTimer", startSleepTimer)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("reached end")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
	<form action="/sleepTimer">
	  <label>Sleeptimer duration (minute):
	    <select name="duration" size="1">
		  <option>1</option>
		  <option>5</option>
		  <option>10</option>
		  <option selected>20</option>
		  <option>30</option>
		</select>
		<input type="submit" value="Go">
	  </label>
	</form>
	</html>`
	fmt.Fprintf(w, html)
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
		log.Fatal("cannot  convert  duration to int: %s", err)
	}

	go sleepTimer(time.Duration(duration))
	fmt.Fprintf(w, "Started sleepTimer (duration: %d minutes)", duration)
}

func sleepTimer(duration time.Duration) {
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
	remaining_time := duration * time.Minute
	interval := remaining_time / 10
	for remaining_time > 0 && brightness > 25 {
		log.Println("Sleeping", interval)
		time.Sleep(interval)
		brightness = brightness - 25
		nk.Set(&hue.State{
			Brightness: brightness,
		})
		log.Println("Decreased brightness to:", nk.State.Brightness, "expected:", brightness)
		remaining_time = remaining_time - interval
		log.Println("Remaining Time:", remaining_time)
	}
	log.Println("Turning light off")
	nk.Off()
}
