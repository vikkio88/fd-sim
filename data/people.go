package data

import (
	"fdsim/enums"
	l "fdsim/locales"
)

func GetNames(country enums.Country) []string {
	switch country {
	case enums.IT:
		return l.ItalianNames
	case enums.EN:
		return l.EnglishNames
	case enums.FR:
		return l.FrenchNames
	case enums.ES:
		return l.SpanishNames
	case enums.DE:
		return l.GermanNames
	}

	return l.ItalianNames
}

func GetSurnames(country enums.Country) []string {
	switch country {
	case enums.IT:
		return l.ItalianSurnames
	case enums.EN:
		return l.EnglishSurnames
	case enums.FR:
		return l.FrenchSurnames
	case enums.ES:
		return l.SpanishSurnames
	case enums.DE:
		return l.GermanSurnames
	}

	return l.ItalianSurnames
}
