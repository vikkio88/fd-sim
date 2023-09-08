package services

import (
	"encoding/json"
	d "fdsim/db"
	"fdsim/enums"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
	"time"
)

func InstantDecisionCancelOffer(offer *models.Offer, db d.IDb) {
	db.MarketR().DeleteOffer(offer)

	//TODO: maybe influence relationships
}

func InstantDecisionBidForAPlayer(params models.EventParams, game *models.Game, db d.IDb) {
	teamId := params.TeamId
	playerId := params.PlayerId
	offerVal := params.ValueF
	fdTeamId := params.FdTeamId

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
		LastUpdate:     game.Date,
	})

	waitingDays := time.Duration(rng.UInt(0, 2))

	var ev d.DbEventDto
	if rng.Chance(chance) {
		ev = d.NewDbEventDto(d.DbEvTeamAcceptedOffer, game.BaseCountry, "", params, game.Date.Add(enums.A_day*waitingDays))
	} else {
		ev = d.NewDbEventDto(d.DbEvTeamRefusedOffer, game.BaseCountry, "", params, game.Date.Add(enums.A_day*waitingDays))
	}

	db.GameR().StoreEvent(ev)

}

func InstantDecisionOfferedContractToAPlayer(params models.EventParams, game *models.Game, db d.IDb) {
	playerId := params.PlayerId
	yContract := params.ValueInt
	wageVal := params.ValueF
	fdTeamId := params.FdTeamId

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
			LastUpdate:     game.Date,
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
		ev = d.NewDbEventDto(d.DbEvPlayerAcceptedContract, game.BaseCountry, "", params, game.Date.Add(enums.A_day*waitingDays))
	} else {
		ev = d.NewDbEventDto(d.DbEvPlayerRefusedContract, game.BaseCountry, "", params, game.Date.Add(enums.A_day*waitingDays))
	}

	db.GameR().StoreEvent(ev)
}

func InstantDecisionConfirmInTransfer(offer *models.Offer, game *models.Game, db d.IDb) {
	//TODO: calculate the Transfer date
	transferDate := game.Date
	offer.TransferDate = transferDate.Add(enums.A_day * 7)

	db.MarketR().SaveOffer(offer)

	ep := models.EP()
	ep.PlayerId = offer.Player.Id
	ep.PlayerName = offer.Player.String()
	ep.FdTeamId = offer.OfferingTeam.Id
	ep.FdTeamName = offer.OfferingTeam.Name

	res, _ := json.Marshal(offer)
	ofjson := string(res)

	ev := d.NewDbEventDto(
		d.DbEvTransferHappening,
		game.BaseCountry,
		ofjson,
		ep,
		offer.TransferDate,
	)

	db.GameR().StoreEvent(ev)
}
