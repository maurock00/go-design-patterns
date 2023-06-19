package main

import (
	"fmt"
	"sync"
)

// Country Mock Database
type Country struct {
	name       string
	code       string
	population int
}

type DatabaseConnector interface {
	CalculatePopulation(country string) int
}

var countries = []Country{
	{"Ecuador", "ECU", 150000000},
	{"Mexico", "MX", 60000000},
	{"Colombia", "CO", 9898348433},
}

type singletonDatabase struct {
	countries []Country
}

func (db *singletonDatabase) CalculatePopulation(country string) int {
	var response int
	for i := range db.countries {
		if db.countries[i].name == country {
			response = db.countries[i].population
		}
	}

	if response > 0 {
		return response
	}

	return 0
}

func SumPopulation(countries []string) int {
	response := 0
	for _, country := range countries {
		response += GetSingletonDatabase().CalculatePopulation(country)
	}

	return response
}

func SumPopulationEx(db DatabaseConnector, countries []string) int {
	response := 0
	for _, country := range countries {
		response += db.CalculatePopulation(country)
	}

	return response
}

type DummieDatabase struct {
}

func (dd *DummieDatabase) CalculatePopulation(country string) int {
	return 100
}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		instance = &singletonDatabase{countries: countries}
		fmt.Println("SOLO SE EJECUTA UNA VEZ")
	})

	return instance
}

func main() {
	singletonInstance := GetSingletonDatabase()
	fmt.Println(singletonInstance.CalculatePopulation("Mexico"))

	singletonInstance2 := GetSingletonDatabase()
	fmt.Println(singletonInstance2.CalculatePopulation("Ecuador"))

	//singletonInstance3 := GetSingletonDatabase()
	//fmt.Println(singletonInstance3.CalculatePopulation("Chile"))

	fmt.Println(SumPopulation([]string{"Mexico", "Ecuador"}))

	fmt.Println(SumPopulationEx(singletonInstance2, []string{"Mexico", "Ecuador"}))

	fmt.Println(SumPopulationEx(&DummieDatabase{}, []string{"Mexico", "Ecuador", "XXXX"}))

}
