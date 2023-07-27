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
type PlayerDetailed struct {
	Player
	History  []*PHistoryRow
	Team     *TPH
	Awards   []Award
	Trophies []Trophy
}

type PlayerHistorical struct {
	Id      string
	Name    string
	Surname string
	Team    *TPH
	Played  int
	Goals   int
	Score   float64
}

func NewPlayerHistoricalFromStatRowPH(row *StatRowPH) *PlayerHistorical {
	return &PlayerHistorical{
		Id:      row.StatRow.Player.Id,
		Name:    row.StatRow.Player.Name,
		Surname: row.StatRow.Player.Surname,
		Team:    row.Team,
		Played:  row.Played,
		Goals:   row.Goals,
		Score:   row.Score,
	}
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

func (p *Player) StringShort() string {
	return fmt.Sprintf("%c. %s", p.Name[0], p.Surname)
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

type RetiredPlayer struct {
	Id      string
	Name    string
	Surname string
	Country enums.Country
	Age     int
	Role    Role

	History     []*PHistoryRow
	Awards      []Award
	Trophies    []Trophy
	YearRetired int
}

func (p *RetiredPlayer) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}
