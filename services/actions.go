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
	ValueInt  *int
	ValueInt2 *int
	ValueF    *float64
	ValueF1   *float64
	Other     *any
}

func MakeActionableFromType(at models.ActionType, date time.Time, params ActionParameter) *models.Actionable {
	switch at {
	case models.ActionRespondContract:
		var yn bool
		return models.NewActionable("Contract Offer", models.Choosable{
			ActionType: at,
			YN:         &yn,
			ValueInt:   params.ValueInt,
			ValueF:     params.ValueF,
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

// can return nil
func ParseDecision(date time.Time, decision *models.Choosable) *Event {
	switch decision.ActionType {
	case models.ActionRespondContract:
		{
			// If decided Yes
			if *decision.YN {
				return ContractAccepted.Event(date, EventParams{
					TeamId1:  *decision.TeamId,
					Label1:   *decision.Label,
					valueInt: *decision.ValueInt,
					valueF:   *decision.ValueF,
				})
			}

			// if refused only reset the flag
			resetFlag := NewEmptyEvent()
			resetFlag.TriggerFlags = func(f models.Flags) models.Flags {
				f.HasAContractOffer = false
				return f
			}
			return resetFlag
		}
	//Testing action
	case models.ActionTest:
		{
			// If decided Yes
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
