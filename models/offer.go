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

func (offer Offer) Stage() OfferStage {
	hasContractOffer := offer.YContract != nil && offer.WageValue != nil

	// Free Agent and Contract offered
	if offer.IsFreeAgent &&
		hasContractOffer &&
		!offer.PlayerAccepted {
		return OfstContractOffered
	}

	// Free Agent and Accepted
	if offer.IsFreeAgent &&
		hasContractOffer &&
		offer.PlayerAccepted {
		return OfstReady
	}

	// Not Free Agent and team has not Accepted
	if !offer.TeamAccepted {
		return OfstOffered
	}

	// Team Accepted, Player hasnt Accepted and No Contract offer
	if offer.TeamAccepted &&
		!hasContractOffer &&
		!offer.PlayerAccepted {
		return OfstTeamAccepted
	}

	if offer.TeamAccepted &&
		hasContractOffer &&
		!offer.PlayerAccepted {
		return OfstContractOffered
	}

	if offer.TeamAccepted && offer.PlayerAccepted {
		return OfstReadyTP
	}

	return OfstNone
}

type OfferStage uint8

const (
	OfstNone OfferStage = iota
	// An Offer has been made to the Team
	OfstOffered
	OfstTeamAccepted
	// An Offer has been made to the Player
	OfstContractOffered
	// Ready Team and Player
	OfstReadyTP

	// Ready Player (When Free agent)
	OfstReady
	OfstFinalised
)

func (o OfferStage) String() string {
	switch o {
	case OfstOffered:
		return "Offer sent"
	case OfstTeamAccepted:
		return "Team Accepted"
	case OfstContractOffered:
		return "Contract Offered"
	case OfstReady:
	case OfstReadyTP:
		return "Ready"
	}

	return "-"
}
