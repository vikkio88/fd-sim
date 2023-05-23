package models

import (
	"fdsim/enums"
	"fmt"

	"github.com/oklog/ulid/v2"
)

const coachInMemoryId = "cmId"

func coachIdGenerator() string {
	return fmt.Sprintf("%s_%s", coachInMemoryId, ulid.Make())
}

type Coach struct {
	idable
	Person
	Module Module
	skillable
}

func NewCoach(name, surname string, age int, country enums.Country, module Module) Coach {
	return Coach{
		idable: NewIdable(coachIdGenerator()),
		Person: Person{
			Name:    name,
			Surname: surname,
			Age:     age,
			Country: country,
		},
		Module: module,
	}
}
