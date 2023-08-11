package models

import (
	"fdsim/utils"
	"time"
)

type Offer struct {
	Player       PNPH
	Team         *TPH
	OfferingTeam TPH

	OfferDate time.Time
	BidValue  *utils.Money
	WageValue *utils.Money
	YContract *int

	IsFreeAgent     bool
	TeamAccepted    bool
	PlayerAccepted  bool
	MoneyTransfered bool
	TransferDate    time.Time
}

func NewOffer(
	player *PNPH,
	offeringTeam *TPH,
	teamAccepted, playerAccepted, moneyTransfered bool,
	offerDate, transferDate time.Time,
) *Offer {
	return &Offer{
		Player:       *player,
		OfferingTeam: *offeringTeam,
		OfferDate:    offerDate,

		TeamAccepted:    teamAccepted,
		PlayerAccepted:  playerAccepted,
		MoneyTransfered: moneyTransfered,
		TransferDate:    transferDate,
	}
}
