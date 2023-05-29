package data

import "fdsim/enums"

func GetTeamNamePattern(country enums.Country) []string {
	switch country {
	case enums.IT:
		return italianTeams
	case enums.EN:
		return englishTeams
	case enums.FR:
		return frenchTeams
	case enums.ES:
		return spanishTeams
	case enums.DE:
		return germanTeams
	}

	return italianTeams
}

var italianTeams = []string{
	"%s Calcio", "%s FC", "AC %s", "AS %s", "%s", "%s Sport", "Città di %s",
}

var spanishTeams = []string{
	"%s CF", "%s FC", "%s United", "Real %s", "Deportivo %s", "Atlético %s", "Sporting %s", "Racing %s", "UD %s", "%s Athletic", "%s",
}

var frenchTeams = []string{
	"Olympique %s", "FC %s", "AS %s", "Stade %s", "%s United", "Racing %s", "ES %s", "Girondins de %s", "OGC %s", "SM Caen %s", "%s",
}

var germanTeams = []string{
	"%s FC", "FC %s", "VfL %s", "Borussia %s", "SV %s", "Eintracht %s", "Hertha %s", "SC %s", "FC Union %s", "1. FC %s", "%s",
}

var englishTeams = []string{
	"AFC %s", "FC %s", "%s United", "%s City", "%s Rovers", "%s Town", "%s Athletic", "%s Albion", "%s Wanderers", "Real %s", "%s Football", "%s",
}
