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

func (ctrl *Controller) Brightness(name string) (uint8, error) {
	ls, err := ctrl.conn.Lights().Get(name)
	if err != nil {
		return 0, err
	}
	return ls.State.Brightness, nil
}

func (ctrl *Controller) SetBrightness(name string, brightness uint8) error {
	ls, err := ctrl.conn.Lights().Get(name)
	if err != nil {
		return err
	}

	return ls.Set(&hue.State{
		Brightness: brightness,
	})
}

func (ctrl *Controller) TurnOffAndRestoreBrightness(name string, brightness uint8) error {
	ls, err := ctrl.conn.Lights().Get(name)
	if err != nil {
		return err
	}

	// use a high transition time (10 seconds) in order to set brightness back to its original value
	// without actually raising the brightness (and wake me up again)
	// it has to be done this way, because brightness cannot be set when the light is turned off
	if err := ls.Set(&hue.State{
		Brightness:     brightness,
		TransitionTime: 100,
	}); err != nil {
		return err
	}

	return ls.Off()
}

func (ctrl *Controller) On(name string) error {
	ls, err := ctrl.conn.Lights().Get(name)
	if err != nil {
		return err
	}

	return ls.On()
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
