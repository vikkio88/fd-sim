package data

import "fdsim/enums"

const (
	ItalyNewspaper   = "Gazzetta dello Sport"
	FranceNewspaper  = "L'Équipe"
	EnglandNewspaper = "BBC Sport"
	SpainNewspaper   = "Marca"
	GermanyNewspaper = "Bild"
)

func GetNewspaper(country enums.Country) string {
	switch country {
	case enums.IT:
		return ItalyNewspaper
	case enums.EN:
		return EnglandNewspaper
	case enums.FR:
		return FranceNewspaper
	case enums.DE:
		return GermanyNewspaper
	case enums.ES:
		return SpainNewspaper
	default:
		return "Unknown Country"
	}
}
