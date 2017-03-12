package main

import (
	"fmt"
	"log"

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
	nk, err := b.Lights().Get("Nachtkaestchen")
	if err != nil {
		log.Fatal(err)
	}
	nk.Toggle()
	//	fmt.Println(nk.State.Brightness)
	fmt.Println("reached end")
}

func startSleepTimer() (_, err) {
}
