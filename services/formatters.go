package services

import (
	"fdsim/data"
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"fmt"
	"strings"
)

// Utils formatters
func emailAddrFromTeamName(teamName string, department string) string {
	website := cleanNameSpaces(teamName, "")
	return fmt.Sprintf("hr@%s.com", website)
}

func getPlayerEmail(playerName string, country enums.Country) string {
	return fmt.Sprintf("%s@%s%s", cleanNameSpaces(playerName, "."), generators.EmailDomains.One(), data.GetDomain(country))
}

func cleanNameSpaces(name string, replace string) string {
	result := strings.ToLower(name)
	result = strings.ReplaceAll(result, " ", replace)
	return result
}

// Link Generators
func teamLink(name, id string) models.Link {
	return models.NewLink(name, enums.TeamDetails, &id)
}

func playerLink(name, id string) models.Link {
	return models.NewLink(name, enums.PlayerDetails, &id)
}

func playerSubLink(name, id string, tabIndex int) models.Link {
	return models.NewLinkSubTab(name, enums.PlayerDetails, &id, &tabIndex)
}
