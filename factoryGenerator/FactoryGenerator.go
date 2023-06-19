package main

import "fmt"

type Employee struct {
	name, position string
	salary         int
}

func EmployeeFactory(position string, salary int) func(name string) *Employee {
	empl := Employee{position: position, salary: salary}
	return func(name string) *Employee {
		empl.name = name

		return &empl
	}
}

func main() {
	developerFactory := EmployeeFactory("developer", 100)
	tlFactory := EmployeeFactory("techlead", 10000)

	dev := developerFactory("Joaquin")
	tl := tlFactory("Fernando")

	fmt.Println(dev)
	fmt.Println(tl)
}
