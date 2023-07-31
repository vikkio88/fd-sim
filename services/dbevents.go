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
	case db.DbEvPlRetiredFdTeam:
		return dbEvRetiredPlayers(dbe, event, params)
	case db.DbEvPlayersSkillChanged:
		return dbEvSkillChangePlayers(dbe, event, params)
	case db.DbEvPlayerLeftFreeAgent:
		return dbEvFreeAgent(dbe, event, params)
	case db.DbEvIndividualAwards:
		return dbEvIndividualAwards(dbe, event, params)
	}

	return nil
}
