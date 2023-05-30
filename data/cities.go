package data

import (
	"fdsim/enums"
	l "fdsim/locales"
)

func GetCities(country enums.Country) []string {
	switch country {
	case enums.IT:
		return l.ItalianCities
	case enums.EN:
		return l.EnglishCities
	case enums.FR:
		return l.FrenchCities
	case enums.ES:
		return l.SpanishCities
	case enums.DE:
		return l.GermanCities
	}

	return l.ItalianCities
}
