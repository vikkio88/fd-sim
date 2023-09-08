package services

import (
	"fdsim/models"
	"time"
)

// can return nil
func ParseDecision(date time.Time, decision *models.Choosable) *Event {
	switch decision.ActionType {
	case models.ActionRespondContract:
		return decisionRespondedToContractOffer(decision, date)

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
