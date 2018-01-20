package controller

import (
	"log"

	"github.com/gbbr/hue"
)

type Controller struct {
	conn *hue.Bridge
}

func NewController() (*Controller, error) {
	a, err := auth()
	if err != nil {
		log.Println("Authentication with Hue Bridge failed: %s", err)
		return nil, err
	}
	return &Controller{
		conn: a,
	}, nil

}

func (ctrl *Controller) AvailableLights() ([]*hue.Light, error) {
	lightsService := ctrl.conn.Lights()
	availableLights, err := lightsService.List()
	if err != nil {
		return nil, err
	}
	return availableLights, nil

}

func auth() (*hue.Bridge, error) {
	b, err := hue.Discover()
	if err != nil {
		log.Println("discovery failed")
		return nil, err
	}
	if !b.IsPaired() {
		log.Println("program not paired with bridge, pairing now")
		// link button must be pressed before calling
		if err := b.Pair(); err != nil {
			log.Println("pairing failed")
			return nil, err
		}
	}
	return b, nil
}
