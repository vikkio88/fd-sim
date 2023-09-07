package services

import (
	"fdsim/db"
	"fdsim/models"
	"fmt"
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

func decisionConfirmInTransfer(decision *models.Choosable, date time.Time) *Event {
	playerId := decision.Params.PlayerId
	fdTeamId := decision.Params.FdTeamId

	//TODO: calculate the Transfer date
	transferDate := date

	ev := NewEmptyEvent()

	// ev.TriggerEmail = models.NewEmail()
	// ev.TriggerNews = models.NewNews()

	ev.TriggerChanges = func(game *models.Game, db db.IDb) {
		o, ok := db.MarketR().GetOffersByPlayerTeamId(playerId, fdTeamId)
		if !ok {
			return
		}

		//TODO: Add ACCEPTED
		o.LastUpdate = date
		o.TransferDate = transferDate

		fmt.Println(o)

	}

	return ev

}
