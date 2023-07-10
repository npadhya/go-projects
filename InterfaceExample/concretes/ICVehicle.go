package concretes

import (
	"errors"
	"log"
)

type ICVehicle struct {
	running bool
}

func NewICVehicle() *ICVehicle {
	return &ICVehicle{
		running: false,
	}
}

func (c *ICVehicle) Start() (bool, error) {
	log.Println("Starting the vehicle")
	if c.running {
		return c.running, errors.New("vehicle already running")
	}
	c.running = true
	return c.running, nil
}

func (c *ICVehicle) Stop() (bool, error) {
	log.Println("Stopping the vehicle")
	if !c.running {
		return c.running, errors.New("vehicle already stopped")
	}
	c.running = false
	return c.running, nil
}

func (c *ICVehicle) IsRunning() bool {
	log.Println("Getting the vehicle status")
	return c.running
}

func (c *ICVehicle) DoRandom() {
	log.Println("ICVehicle doing random stuff")
}
