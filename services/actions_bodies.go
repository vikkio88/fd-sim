package services

import (
	"fdsim/models"
	"time"
)

func decisionRespondedToContractOffer(decision *models.Choosable, date time.Time) *Event {
	if *decision.YN {
		return ContractAccepted.Event(date, decision.Params)
	}

	// TODO: If player responded NO maybe inhibit team to offer again: relationship
	resetFlag := NewEmptyEvent()
	resetFlag.TriggerFlags = func(f models.Flags) models.Flags {
		f.HasAContractOffer = false
		return f
	}
	return resetFlag
}
