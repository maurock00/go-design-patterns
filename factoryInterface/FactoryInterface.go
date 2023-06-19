package main

import (
	"fmt"
	"strconv"
)

type Vehicle interface {
	RenderVehicle()
}

type vehicle struct {
	tiresNumber int
	color       string
}

type motorcycle struct {
	tiresNumber int
	color       string
}

func (v *vehicle) RenderVehicle() {
	fmt.Println("This vehicle has " + strconv.Itoa(v.tiresNumber) + " tires and is " + v.color)
}

func (v *motorcycle) RenderVehicle() {
	fmt.Println("This MOTORCYLCE has " + strconv.Itoa(v.tiresNumber) + " tires and is " + v.color)
}

func CreateVehicle(tires int, color string) Vehicle {
	if tires == 2 {
		return &motorcycle{tiresNumber: tires, color: color}
	}
	return &vehicle{tiresNumber: tires, color: color}
}

func main() {
	newV := CreateVehicle(2, "red")
	newV.RenderVehicle()

	newV2 := CreateVehicle(100, "green")
	newV2.RenderVehicle()
}
