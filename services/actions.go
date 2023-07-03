package services

import (
	"fdsim/models"
	"time"
)

func MakeActionableFromType(at models.ActionType, date time.Time, params models.EventParams) *models.Actionable {
	switch at {
	case models.ActionRespondContract:
		var yn bool
		return models.NewActionable("Contract Offer", models.Choosable{
			ActionType: at,
			YN:         &yn,
			Params:     params,
		},
			at,
		)

		//TODO: Remove Testing Action
	case models.ActionTest:
		var yn bool
		return models.NewActionable("Testing Actionables", models.Choosable{
			ActionType: at,
			YN:         &yn,
			Params:     params,
		},
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
				return ContractAccepted.Event(date, decision.Params)
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
				return TestingActionYes.Event(date, decision.Params)
			}

			return TestingActionNo.Event(date, decision.Params)
		}

	}
	return nil
}
