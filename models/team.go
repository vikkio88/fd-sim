package models

import (
	"fdsim/enums"
	"fmt"

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

	//TODO: add familiarity with a module
}

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

func (t *Team) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.Country)
}

func (t *Team) StringShort() string {
	return fmt.Sprintf("%s", t.Name)
}

func (t *Team) Lineup() *Lineup {
	module := M442
	if t.Coach != nil {
		module = t.Coach.Module
	}

	return t.Roster.Lineup(module)
}

func (t *Team) PH() TPH {
	return TPH{
		Id:   t.Id,
		Name: t.Name,
	}
}
