package main

import (
	"testing"
)


func TestMain(t *testing.T)  {
	
	t.Run("processTruck", func(t *testing.T){

		t.Run("should load and unload a truck cargo", func(t *testing.T){

			normalTruck := &NormalTruck{id: "1",cargo: 42}
	electricTruck := &ElectricTruck{id: "2"}

	
	err := processTruck(normalTruck)
	if err != nil {
		t.Fatalf("Error processing truck: %s", err)

	}

	err = processTruck(electricTruck)
	if err != nil {
		t.Fatalf("Error processing electric truck: %s", err)
	}
	if normalTruck.cargo != 0 {
		t.Fatal("Normal truck cargo should be 0")
	}
			
		})
	})
}