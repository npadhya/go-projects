package concretes

import (
	"errors"
	"log"
)

type ElectricVehicle struct {
	running bool
}

func NewElectricVehicle() *ElectricVehicle {
	return &ElectricVehicle{
		running: false,
	}
}

func (c *ElectricVehicle) Start() (bool, error) {
	log.Println("Starting the vehicle")
	if c.running {
		return c.running, errors.New("vehicle already running")
	}
	c.running = true
	return c.running, nil
}

func (c *ElectricVehicle) Stop() (bool, error) {
	log.Println("Stopping the vehicle")
	if !c.running {
		return c.running, errors.New("vehicle already stopped")
	}
	c.running = false
	return c.running, nil
}

func (c *ElectricVehicle) IsRunning() bool {
	log.Println("Getting the vehicle status")
	return c.running
}

func (c *ElectricVehicle) DoRandom() {
	log.Println("ElectricVehicle doing random stuff")
}
