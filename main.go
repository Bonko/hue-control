package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gbbr/hue"
)

func main() {
	b, err := hue.Discover()
	if err != nil {
		fmt.Println("discovery failed")
		log.Fatal(err)
	}
	if !b.IsPaired() {
		// link button must be pressed before calling
		if err := b.Pair(); err != nil {
			log.Fatal(err)
		}
	}
	sleepTimer(b)
	fmt.Println("reached end")
}

func sleepTimer(b *hue.Bridge) {
	nk, err := b.Lights().Get("Nachtkaestchen")
	if err != nil {
		log.Fatal(err)
	}
	brightness := uint8(255)
	fmt.Println("Turning light on")
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
