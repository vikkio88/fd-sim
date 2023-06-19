package models

import (
	"fdsim/enums"
	"fdsim/utils"
	"fmt"

	"github.com/oklog/ulid/v2"
)

const playerInMemoryId = "pmId"

func playerIdGenerator() string {
	return fmt.Sprintf("%s_%s", playerInMemoryId, ulid.Make())
}

type Player struct {
	Idable
	Person
	Role  Role
	Value utils.Money
	skillable
}
type PlayerWithTeam struct {
	Player
	Team *TPH
}

func NewPlayer(name, surname string, age int, country enums.Country, role Role) Player {
	return Player{
		Idable: NewIdable(playerIdGenerator()),
		Person: Person{
			Name:    name,
			Surname: surname,
			Age:     age,
			Country: country,
		},
		Role: role,
		//TODO: add familiarity with a module
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

// get placeholder
func (p *Player) PH() PPH {
	return PPH{
		Id:  p.Id,
		sPH: p.skillable.PH(),
		Age: p.Age,
	}
}

// Player Placeholder name
type PNPH struct {
	Id      string
	Name    string
	Surname string
}

func (p *PNPH) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

// Player Placeholder without names
type PPH struct {
	sPH
	Id  string
	Age int
	// TODO: track injuries so we know whether can be choose or not for lineup
}

func NewRolePPHMap() map[Role][]PPH {
	result := map[Role][]PPH{}
	for _, role := range AllPlayerRoles() {
		result[role] = []PPH{}
	}

	return result
}
