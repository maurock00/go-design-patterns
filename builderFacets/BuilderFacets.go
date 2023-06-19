package main

import "fmt"

type AddressInfo struct {
	City   string
	Street string
}

type JobInfo struct {
	Company string
	Salary  float32
}

type PersonV2 struct {
	AddressInfo AddressInfo
	JobInfo     JobInfo
}

type PersonBuilder struct {
	Person *PersonV2
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&PersonV2{}}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (ab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	ab.Person.AddressInfo.Street = city
	return ab
}

func (ab *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	ab.Person.AddressInfo.Street = street
	return ab
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (jb *PersonJobBuilder) At(company string) *PersonJobBuilder {
	jb.Person.JobInfo.Company = company
	return jb
}

func (jb *PersonJobBuilder) Earn(salary float32) *PersonJobBuilder {
	jb.Person.JobInfo.Salary = salary
	return jb
}

func (b *PersonBuilder) Build() *PersonV2 {
	return b.Person
}

func main() {
	personBuilder := NewPersonBuilder()

	newPerson := personBuilder.
		Works().At("Kushki").Earn(100.20).
		Lives().At("America y asunción").In("Quito").
		Build()

	fmt.Println(newPerson)
}
