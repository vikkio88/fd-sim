package models

import (
	"fdsim/enums"
	"fmt"

	"github.com/oklog/ulid/v2"
)

const teamInMemoryId = "tmId"

func teamIdGenerator() string {
	return fmt.Sprintf("%s_%s", playerInMemoryId, ulid.Make())
}

type Team struct {
	idable
	Name    string
	City    string
	Country enums.Country
	Roster  *Roster
	Coach   *Coach

	//TODO: add familiarity with a module
}

func NewTeam(name, city string, country enums.Country) Team {
	return Team{
		idable:  NewIdable(teamIdGenerator()),
		Name:    name,
		City:    city,
		Country: country,
		Roster:  NewRoster(),
	}
}
