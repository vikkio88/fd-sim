package services

import (
	d "fdsim/db"
	"fdsim/enums"
	"fdsim/libs"
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
	teamId := decision.Params.TeamId
	playerId := decision.Params.PlayerId
	offerVal := decision.Params.ValueF
	event := NewEmptyEvent()

	event.TriggerChanges = func(game *models.Game, db d.IDb) {
		t, ok := db.TeamR().ById(teamId)
		if !ok {
			return
		}

		player, ok := t.Roster.Player(playerId)
		if !ok {
			return
		}

		rng := libs.NewRngAutoSeeded()

		//TODO: check if player is on the market

		chance := 50
		if offerVal >= player.Value.Value() {
			chance += 30
		} else {
			chance -= 10
		}

		//TODO: check if team is skint/ need that player

		var ev d.DbEventDto
		if rng.ChanceI(chance) {
			ev = d.NewDbEventDto(d.DbEvTeamAcceptedOffer, game.BaseCountry, "", decision.Params, game.Date.Add(enums.A_day*2))
		} else {
			ev = d.NewDbEventDto(d.DbEvTeamRefusedOffer, game.BaseCountry, "", decision.Params, game.Date.Add(enums.A_day*2))
		}

		db.GameR().StoreEvent(ev)

	}
	return event
}
