package services

import (
	"fdsim/models"
	"time"
)

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
