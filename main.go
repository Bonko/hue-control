package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gbbr/hue"
)

func main() {
	fmt.Println("listening..")
	//	sleepTimer()
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/sleepTimer", sleepTimer)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("reached end")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "bla")
}

func auth() *hue.Bridge {
	b, err := hue.Discover()
	if err != nil {
		fmt.Println("discovery failed")
		log.Fatal(err)
	}
	if !b.IsPaired() {
		fmt.Println("program not paired with bridge")
		// link button must be pressed before calling
		if err := b.Pair(); err != nil {
			fmt.Println("pairing failed")
			log.Fatal(err)
		}
	}
	return b
}

func sleepTimer(w http.ResponseWriter, r *http.Request) {
	b := auth()
	nk, err := b.Lights().Get("Nachtkaestchen")
	if err != nil {
		log.Fatal(err)
	}
	brightness := uint8(255)
	fmt.Fprintf(w, "Turning light on")
	nk.On()
	nk.Set(&hue.State{
		Brightness: brightness,
	})
	remaining_time := 20 * time.Minute
	for remaining_time > 0 && brightness > 25 {
		interval := 2 * time.Minute
		fmt.Println("Sleeping", interval)
		time.Sleep(interval)
		brightness = brightness - 25
		fmt.Println("Decreasing brightness to", brightness)
		nk.Set(&hue.State{
			Brightness: brightness,
		})
		fmt.Println(nk.State.Brightness)
		remaining_time = remaining_time - interval
		fmt.Println("Remaining Time:", remaining_time)
	}
	nk.Off()
}
