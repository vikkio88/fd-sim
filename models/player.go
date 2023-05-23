package models

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

const playerInMemoryId = "pmId"

func playerIdGenerator() string {
	return fmt.Sprintf("%s_%s", playerInMemoryId, ulid.Make())
}

type Player struct {
	idable
	Person
	Role Role
	skillable
}

func NewPlayer(name, surname string, age int, role Role) Player {
	return Player{
		idable: NewIdable(playerIdGenerator()),
		Person: Person{
			Name:    name,
			Surname: surname,
			Age:     age,
		},
		Role: role,
	}
}
