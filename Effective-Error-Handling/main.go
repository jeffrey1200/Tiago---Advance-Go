package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

func (t *Truck) LoadCargo() error {
	return ErrTruckNotFound
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {

	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return ErrNotImplemented
}

func main() {

	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		// err := processTruck(truck)
		// if err != nil {
		// 	log.Fatalf("Error processing truck: %s",err)
		// }
		if err := processTruck(truck); err != nil {

			if errors.Is(err, ErrTruckNotFound) {
				log.Fatal("TRUE")
			}

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
			log.Fatalf("Error processing truck: %s", err)
		}
	}
}
