package main

import "fmt"

type car struct {
	carType string
	fuelin  float32
}

func (c car) carFuel() float32 {
	result := c.fuelin / 1.5
	return result
}

func main() {
	c := car{
		carType: "Suv",
		fuelin:  10.5,
	}
	fmt.Println("TypeCar: "+c.carType+"\nDistance : ", c.carFuel())
}
