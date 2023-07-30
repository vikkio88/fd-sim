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
				emailAddrFromTeamName(params.TeamName, "hr"),
				fmt.Sprintf("Goodbye! %d Players retired.", len(retired)),
				body,
				dbe.TriggerDate,
				links,
			)
			return event
		}
	case db.DbEvPlayersSkillChanged:
		{
			var ps [2][]*models.PNPHVals
			json.Unmarshal([]byte(dbe.Payload), &ps)
			improved := ps[0]
			worsened := ps[1]

			links := []models.Link{}

			body := "Few players this year showcased a significant change in their skills."
			if len(improved) > 0 {
				body += "\nThe following players improved quite a lot:"
				for _, v := range improved {
					links = append(links, models.NewLink(fmt.Sprintf("%s : +%d%%", v.String(), v.ValueI), enums.PlayerDetails, &v.Id))
					body += fmt.Sprintf(" %s", conf.LinkBodyPH)
				}
			}
			if len(worsened) > 0 {
				body += "\nThose players skills appear to be worse:"
				for _, v := range worsened {
					links = append(links, models.NewLink(fmt.Sprintf("%s : -%d%%", v.String(), v.ValueI), enums.PlayerDetails, &v.Id))
					body += fmt.Sprintf(" %s", conf.LinkBodyPH)
				}
			}

			event.TriggerEmail = models.NewEmail(
				emailAddrFromTeamName(params.TeamName, "training"),
				"Report of your player skill change.",
				body,
				dbe.TriggerDate,
				links,
			)

			return event
		}
	}

	return nil
}
