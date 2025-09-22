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

	// trucks := []NormalTruck{
	// 	{id: "Truck-1"},
	// 	{id: "Truck-2"},
	// 	{id: "Truck-3"},
	// }

	// eTrucks := []ElectricTruck{
	// 	{id: "electric-truck-1"},
	// }
	normalTruck := &NormalTruck{id: "1"}
	electricTruck := &ElectricTruck{id: "2"}
	err := processTruck(normalTruck)
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)

	}

	err = processTruck(electricTruck)
	if err != nil {
		log.Fatalf("Error processing electric truck: %s", err)
	}

	log.Println(normalTruck.cargo)
	log.Println(electricTruck.battery)
	// for _, truck := range trucks {
	// fmt.Printf("Truck %s arrived.\n", truck.id)
	// err := processTruck(truck)
	// if err != nil {
	// 	log.Fatalf("Error processing truck: %s",err)
	// }
	// if err := processTruck(truck); err != nil {

	// log.Fatalf("Error processing truck: %s", err)

	// if errors.Is(err, ErrTruckNotFound) {
	// 	log.Fatal("TRUE")
	// }

	// switch err {
	// case ErrNotImplemented:
	// 	log.Fatalf("this feature is not implemented yet: %s", err)

	// case ErrTruckNotFound:
	// 	log.Fatalf("error processing truck: %s", err)

	// default:
	// 	log.Printf("Hey you've hit the default case")
	// }

	// if errors.Is(err,ErrNotImplemented) {

	// }
	// if errors.Is(err,ErrTruckNotFound){

	// }
	// }
	// }
}
