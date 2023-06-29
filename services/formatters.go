package services

import (
	"fdsim/enums"
	"fdsim/models"
	"fmt"
	"strings"
)

// Utils formatters
func emailAddrFromTeamName(teamName string) string {
	website := strings.ToLower(teamName)
	website = strings.ReplaceAll(website, " ", "")
	return fmt.Sprintf("hr@%s.com", website)
}

// Link Generators
func teamLink(name, id string) models.Link {
	return models.NewLink(name, enums.TeamDetails, &id)
}
