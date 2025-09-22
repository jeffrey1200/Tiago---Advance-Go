package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery = -1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery += -1
	return nil
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {

	fmt.Printf("Processing truck: %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	t := NormalTruck{cargo: 0}
	fillTruck(&t)

	log.Println(t)
	truckID := 42
	anotherTruckID := &truckID
	fmt.Println(truckID)
	fmt.Println(*anotherTruckID)
	
	truckID = 0;

	fmt.Println(*anotherTruckID)
}

func fillTruck(t *NormalTruck) {
	t.cargo = 100
}