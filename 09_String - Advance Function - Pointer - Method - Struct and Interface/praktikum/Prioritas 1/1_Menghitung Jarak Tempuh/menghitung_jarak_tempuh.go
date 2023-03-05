package main

import (
	"fmt"
)

type Car struct {
	typecar string
	fuelIn  float64
}

func (C Car) Mileage() float64 {
	return C.fuelIn / 1.5
}

func main() {
	CarA := Car{"Civic", 30}

	fmt.Printf("%v\n", CarA)
	fmt.Printf("Jarak yang ditempuh : ")
	fmt.Println(CarA.Mileage())
}
