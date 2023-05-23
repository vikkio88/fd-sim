package models

import "fdsim/enums"

type Team struct {
	Name    string
	City    string
	Country enums.Country
	Roster  *Roster
	Coach   *Coach
}

func NewTeam(name, city string, country enums.Country) Team {
	return Team{
		Name:    name,
		City:    city,
		Country: country,
	}
}
