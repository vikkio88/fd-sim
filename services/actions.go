package services

import (
	"fdsim/models"
	"time"
)

const (
	a_day = time.Duration(1) * time.Hour * 24
)

type ActionParameter struct {
	LeagueId  *string
	TeamId    *string
	TeamId1   *string
	TeamId2   *string
	PlayerId  *string
	PlayerId1 *string
	PlayerId2 *string
	Label     *string
	Label1    *string
	Label2    *string
	Label3    *string
	Value     *float64
	Other     *any
}

func MakeActionable(at models.ActionType, date time.Time, params ActionParameter) *models.Actionable {
	switch at {
	case models.ActionContract:
		var yn bool
		return models.NewActionable("Contract Offer", models.Choosable{
			ActionType: at,
			YN:         &yn,
			Value:      params.Value,
			TeamId:     params.TeamId,
			Label:      params.Label,
		},
			date.Add(2*a_day),
			at,
		)

		//TODO: Remove Testing Action
	case models.ActionTest:
		var yn bool
		return models.NewActionable("Testing Actionables", models.Choosable{
			ActionType: at,
			YN:         &yn,
			TeamId:     params.TeamId,
			Label:      params.Label,
		},
			date.Add(2*a_day),
			at,
		)
	}

	return nil
}

func ParseDecision(date time.Time, decision *models.Choosable) *Event {
	switch decision.ActionType {
	case models.ActionTest:
		{
			if *decision.YN {
				return TestingActionYes.Event(date, EventParams{
					TeamId1: *decision.TeamId,
					Label1:  *decision.Label,
				})
			}

			return TestingActionNo.Event(date, EventParams{
				TeamId1: *decision.TeamId,
				Label1:  *decision.Label,
			})
		}

	}
	return nil
}
