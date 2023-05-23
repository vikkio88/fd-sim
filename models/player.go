package models

type Player struct {
	Person
	Role Role
	skillable
}

func NewPlayer(name, surname string, age int, role Role) Player {
	return Player{
		Person: Person{
			Name:    name,
			Surname: surname,
			Age:     age,
		},
		Role: role,
	}
}
