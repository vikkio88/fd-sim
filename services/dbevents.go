package services

import (
	"encoding/json"
	"fdsim/conf"
	"fdsim/db"
	"fdsim/enums"
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
		{
			var retired []*models.PNPH
			json.Unmarshal([]byte(dbe.Payload), &retired)

			links := make([]models.Link, len(retired))
			body := "The following players retired this year"
			for i, rp := range retired {
				links[i] = models.NewLink(rp.String(), enums.PlayerDetails, &rp.Id)
				body += fmt.Sprintf(" %s", conf.LinkBodyPH)
			}

			event.TriggerEmail = models.NewEmail(
				emailAddrFromTeamName(params.TeamName),
				fmt.Sprintf("%d Players retired.", len(retired)),
				body,
				dbe.TriggerDate,
				links,
			)
			return event
		}
	}

	return nil
}
