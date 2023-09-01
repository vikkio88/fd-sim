package db

import (
	"fdsim/models"
	"fdsim/utils"
	"time"
)

type OfferDto struct {
	PlayerId string
	TeamId   *string

	OfferingTeamId string
	OfferDate      time.Time
	BidValue       *float64
	WageValue      *float64
	YContract      *int

	TeamAccepted    bool
	PlayerAccepted  bool
	MoneyTransfered bool
	TransferDate    time.Time

	Player PlayerDto
	Team   *TeamDto

	OfferingTeam TeamDto
}

func DtoFromOffer(offer *models.Offer) OfferDto {
	o := OfferDto{
		PlayerId:       offer.Player.Id,
		OfferingTeamId: offer.OfferingTeam.Id,
		OfferDate:      offer.OfferDate,

		TeamAccepted:    offer.TeamAccepted,
		PlayerAccepted:  offer.PlayerAccepted,
		MoneyTransfered: offer.MoneyTransfered,
		TransferDate:    offer.TransferDate,
	}

	if offer.BidValue != nil {
		b := offer.BidValue.Value()
		o.BidValue = &b
	}
	if offer.WageValue != nil {
		w := offer.WageValue.Value()
		o.WageValue = &w
	}
	if offer.YContract != nil {
		c := offer.YContract
		o.YContract = c
	}

	if offer.Team != nil {
		o.TeamId = &offer.Team.Id
	}

	return o
}

func (o *OfferDto) Offer() *models.Offer {
	offer := models.NewOffer(
		o.Player.PlayerPH(),
		o.OfferingTeam.TeamPH(),
		o.TeamAccepted, o.PlayerAccepted, o.MoneyTransfered,
		o.OfferDate, o.TransferDate,
	)
	if o.Team != nil {
		offer.Team = o.Team.TeamPH()
	} else {
		// this is maybe not needed
		offer.IsFreeAgent = true
	}

	if o.BidValue != nil {
		b := utils.NewEurosFromF(*o.BidValue)
		offer.BidValue = &b
	}

	if o.WageValue != nil {
		w := utils.NewEurosFromF(*o.WageValue)
		offer.WageValue = &w
	}

	if o.YContract != nil {
		c := o.YContract
		offer.YContract = c
	}

	if o.Team != nil {
		offer.Team = o.Team.TeamPH()
	}

	return offer
}
