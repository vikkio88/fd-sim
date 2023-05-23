package models

type Coach struct {
	Person
	Module Module
	skillable
}

func NewCoach(name, surname string, age int, module Module) Coach {
	return Coach{
		Person: Person{
			Name:    name,
			Surname: surname,
			Age:     age,
		},
		Module: module,
	}
}
