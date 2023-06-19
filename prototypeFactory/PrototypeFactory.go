package main

import "fmt"

type Pet struct {
	Name, Kind string
	Weight     int
}

const (
	DOG = "DOG"
	CAT = "CAT"
)

func NewPet(kind string) *Pet {
	switch kind {
	case CAT:
		return &Pet{"", CAT, 10}
	case DOG:
		return &Pet{"", DOG, 20}
	default:
		panic("No valid kind of animal")
	}
}

func main() {
	dog := NewPet(DOG)
	cat := NewPet(CAT)

	fmt.Println(dog)
	fmt.Println(cat)
	fmt.Println(NewPet("kjassjkakja"))
}
