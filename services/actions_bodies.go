package services

import (
	"fdsim/db"
	"fdsim/models"
	"time"
)

func decisionRespondedToContractOffer(decision *models.Choosable, date time.Time) *Event {
	if *decision.YN {
		return ContractAccepted.Event(date, decision.Params)
	}

	// TODO: If player responded NO maybe inhibit team to offer again
	resetFlag := NewEmptyEvent()
	resetFlag.TriggerFlags = func(f models.Flags) models.Flags {
		f.HasAContractOffer = false
		return f
	}
	return resetFlag
}

func decisionOfferedForAPlayer(decision *models.Choosable, date time.Time) *Event {

	// 	teamId := decision.Params.TeamId
	// 	playerId := decision.Params.PlayerId
	// 	offerVal := decision.Params.ValueF
	event := NewEmptyEvent()

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		// maybe here actually find a way to choose whether offer is ok
	}
	return event
}
