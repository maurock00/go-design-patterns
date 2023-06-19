package main

import "fmt"

type PersonV3 struct {
	name, role string
}

type PersonBuilderAction func(*PersonV3)

type PersonBuilderV3 struct {
	actions []PersonBuilderAction
}

func (b *PersonBuilderV3) Called(name string) *PersonBuilderV3 {
	b.actions = append(b.actions, func(person *PersonV3) {
		person.name = name
	})

	return b
}

func (b *PersonBuilderV3) WorkAs(role string) *PersonBuilderV3 {
	b.actions = append(b.actions, func(person *PersonV3) {
		person.role = role
	})

	return b
}

func (b *PersonBuilderV3) build() *PersonV3 {
	person := PersonV3{}
	for _, action := range b.actions {
		action(&person)
	}

	return &person
}

func printPersonImpl(person *PersonV3) {
	fmt.Println(person)
}

type buildPerson func(*PersonBuilderV3)

func PrintPerson(action buildPerson) {
	builder := PersonBuilderV3{}
	action(&builder)
	printPersonImpl(builder.build())
}

func main() {
	builderFunction := func(p3 *PersonBuilderV3) {
		p3.Called("Maurocker").WorkAs("Architect")
	}

	PrintPerson(builderFunction)
}
