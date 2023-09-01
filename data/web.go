package data

import "fdsim/enums"

var EmailDomains = []string{
	"gmail",
	"yahoo",
	"outlook",
	"hotmail",
	"icloud",
	"mail",
	"protonmail",
	"live",
	"inbox",
}

func GetDomain(country enums.Country) string {
	switch country {
	case enums.IT:
		return ".it"
	case enums.EN:
		return ".co.uk"
	case enums.FR:
		return ".fr"
	case enums.ES:
		return ".es"
	case enums.DE:
		return ".de"
	}

	return ".com"
}
