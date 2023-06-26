package data

import "fdsim/enums"

const (
	ItalianLeague = "Serie A"
	EnglishLeague = "Premier League"
	FrenchLeague  = "Ligue 1"
	GermanLeague  = "Bundesliga"
	SpanishLeague = "La Liga"
)

func GetLeagueName(country enums.Country) string {
	switch country {
	case enums.IT:
		return ItalianLeague
	case enums.EN:
		return EnglishLeague
	case enums.FR:
		return FrenchLeague
	case enums.DE:
		return GermanLeague
	case enums.ES:
		return SpanishLeague
	default:
		return "Unknown Country"
	}
}
