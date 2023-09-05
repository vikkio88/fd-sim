package services

import (
	d "fdsim/db"
	"fdsim/enums"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
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
	fdTeamId := decision.Params.FdTeamId
	event := NewEmptyEvent()

	event.TriggerChanges = func(game *models.Game, db d.IDb) {
		t, ok := db.TeamR().ById(teamId)
		if !ok {
			return
		}
		// getting rng tied in with offer so it will always be the same
		rng := libs.NewRng(int64(offerVal))

		//TODO: check if player is on the market

		chance := t.OfferAcceptanceChance(utils.NewEurosFromF(offerVal), playerId)

		//TODO: check if team is skint/ need that player

		// Persist Offer on Db
		db.MarketR().AddOffer(d.OfferDto{
			PlayerId:       playerId,
			TeamId:         &teamId,
			OfferingTeamId: fdTeamId,
			BidValue:       &offerVal,
			OfferDate:      game.Date,
		})

		waitingDays := time.Duration(rng.UInt(0, 2))

		var ev d.DbEventDto
		if rng.Chance(chance) {
			ev = d.NewDbEventDto(d.DbEvTeamAcceptedOffer, game.BaseCountry, "", decision.Params, game.Date.Add(enums.A_day*waitingDays))
		} else {
			ev = d.NewDbEventDto(d.DbEvTeamRefusedOffer, game.BaseCountry, "", decision.Params, game.Date.Add(enums.A_day*waitingDays))
		}

		db.GameR().StoreEvent(ev)

	}
	return event
}

func decisionOfferedContractToAPlayer(decision *models.Choosable, date time.Time) *Event {
	playerId := decision.Params.PlayerId
	yContract := decision.Params.ValueInt
	wageVal := decision.Params.ValueF
	fdTeamId := decision.Params.FdTeamId
	event := NewEmptyEvent()

	event.TriggerChanges = func(game *models.Game, db d.IDb) {
		p, ok := db.PlayerR().ById(playerId)
		if !ok {
			return
		}
		// getting rng tied in with offer so it will always be the same
		rng := libs.NewRng(int64(wageVal))
		offeredWage := utils.NewEurosFromF(wageVal)
		fdTeam, ok := db.TeamR().ById(fdTeamId)
		if !ok {
			return
		}
		chance := p.WageAcceptanceChance(offeredWage, yContract, fdTeam)

		// get offer if exists (this means it was offered a bid to the team)
		o, exists := p.GetOfferFromTeamId(fdTeamId)
		if !exists {
			// if it doesnt exist this means that is a free agent
			//TODO: I could just check the Team != nil but I want to make sure
			newOffer := d.OfferDto{
				PlayerId:       playerId,
				OfferingTeamId: fdTeamId,
				WageValue:      &wageVal,
				YContract:      &yContract,
				OfferDate:      game.Date,
			}
			db.MarketR().AddOffer(newOffer)
		} else {
			o.WageValue = &offeredWage
			o.YContract = &yContract
			db.MarketR().SaveOffer(o)
		}
		// Persist Offer on Db or Update if is there already

		waitingDays := time.Duration(rng.UInt(0, 2))

		var ev d.DbEventDto
		if rng.Chance(chance) {
			ev = d.NewDbEventDto(d.DbEvPlayerAcceptedContract, game.BaseCountry, "", decision.Params, game.Date.Add(enums.A_day*waitingDays))
		} else {
			ev = d.NewDbEventDto(d.DbEvPlayerRefusedContract, game.BaseCountry, "", decision.Params, game.Date.Add(enums.A_day*waitingDays))
		}

		db.GameR().StoreEvent(ev)

	}
	return event
}
