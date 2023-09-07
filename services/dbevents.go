package services

import (
	"encoding/json"
	"fdsim/db"
	"fdsim/models"
	"fmt"
)

func parseEventParams(ep string) models.EventParams {
	if ep == "" {
		return models.EventParams{}
	}

	var result models.EventParams

	err := json.Unmarshal([]byte(ep), &result)
	if err != nil {
		fmt.Println("error while loading", err)
		return models.EventParams{}
	}

	return result
}

func getEventFromDbEvent(dbe db.DbEventDto) *Event {
	event := NewEvent(dbe.TriggerDate, "db_event")
	params := parseEventParams(dbe.EventParams)

	switch dbe.Type {
	case db.DbEvYoungJoinedFdTeam:
		return dbEvYoungPlayers(dbe, event, params)
	case db.DbEvPlRetiredFdTeam:
		return dbEvRetiredPlayers(dbe, event, params)
	case db.DbEvPlayersSkillChanged:
		return dbEvSkillChangePlayers(dbe, event, params)
	case db.DbEvPlayerLeftFreeAgent:
		return dbEvFreeAgent(dbe, event, params)
	case db.DbEvIndividualAwards:
		return dbEvIndividualAwards(dbe, event, params)
	case db.DbEvTeamAcceptedOffer:
		return dbEvTeamAcceptedOffer(dbe, event, params)
	case db.DbEvTeamRefusedOffer:
		return dbEvTeamRefusedOffer(dbe, event, params)
	case db.DbEvPlayerAcceptedContract:
		return dbEvPlayerAcceptedContract(dbe, event, params)
	case db.DbEvPlayerRefusedContract:
		return dbEvPlayerRefusedContract(dbe, event, params)
	case db.DbEvTransferConfirmed:
		return dbEvTransferConfirmed(dbe, event, params)
	}

	return nil
}
