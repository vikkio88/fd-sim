package models

import (
	"fdsim/enums"
	"fdsim/utils"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

const teamInMemoryId = "tmId"

func teamIdGenerator() string {
	return fmt.Sprintf("%s_%s", teamInMemoryId, ulid.Make())
}

type Team struct {
	Idable
	Name    string
	City    string
	Country enums.Country
	Roster  *Roster
	Coach   *Coach

	Balance       utils.Money
	TransferRatio float64

	//TODO: add familiarity with a module
}

// Team Placeholder
type TPH struct {
	Id   string
	Name string
}

func NewTeam(name, city string, country enums.Country) *Team {
	return &Team{
		Idable:  NewIdable(teamIdGenerator()),
		Name:    name,
		City:    city,
		Country: country,
		Roster:  NewRoster(),
	}
}

func (t *Team) Wages() utils.Money {
	r := t.Roster.Wages()
	return utils.NewEurosFromF(r.Value() + t.Coach.Wage.Value())
}
func (t *Team) TransferBudget() utils.Money {
	val := t.Balance.Value() * t.TransferRatio
	return utils.NewEurosFromF(val)
}

func (t *Team) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.Country)
}

func (t *Team) StringShort() string {
	return fmt.Sprintf("%s", t.Name)
}

func (t *Team) Lineup() *Lineup {
	module := M442
	rngSeed := time.Now().Unix()
	morale := 50
	if t.Coach != nil {
		module = t.Coach.Module
		rngSeed = t.Coach.RngSeed
		morale = t.Coach.Morale.Val()
	}

	return t.Roster.Lineup(module, rngSeed, morale)
}

func (t *Team) PH() TPH {
	return TPH{
		Id:   t.Id,
		Name: t.Name,
	}
}
