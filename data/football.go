package data

import (
	"fdsim/enums"
	l "fdsim/locales"
)

func GetTeamNamePattern(country enums.Country) []string {
	switch country {
	case enums.IT:
		return l.ItalianTeams
	case enums.EN:
		return l.EnglishTeams
	case enums.FR:
		return l.FrenchTeams
	case enums.ES:
		return l.SpanishTeams
	case enums.DE:
		return l.GermanTeams
	}

	return l.ItalianTeams
}
