package main

import (
	"log"
	"testing"
)


func TestMain(t *testing.T)  {
	
	t.Run("processTruck", func(t *testing.T){

		t.Run("should load and unload a truck cargo", func(t *testing.T){

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
			
		})
	})
}