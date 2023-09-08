package db

import (
	"fdsim/models"

	"gorm.io/gorm"
)

type MarketRepo struct {
	g *gorm.DB
}

func NewMarketRepo(g *gorm.DB) *MarketRepo {
	return &MarketRepo{g}
}

func (repo *MarketRepo) Truncate() {
	repo.g.Where("1 = 1").Delete(&OfferDto{})
}

func (repo *MarketRepo) GetTransferMarketInfo() (*models.TransferMarketInfo, bool) {
	var dto transfMkInfoDto
	repo.g.Raw("select td.id as team_id, td.balance, td.transfer_ratio from game_dtos gd left join team_dtos td on gd.team_id = td.id;").Find(&dto)

	if dto.TeamId == "" {
		return nil, false
	}

	return dto.TransferMarketInfo(), true
}

func (repo *MarketRepo) AddOffer(offer OfferDto) {
	repo.g.Create(&offer)
}

func (repo *MarketRepo) SaveOffer(offer *models.Offer) {
	playerId := offer.Player.Id
	offeringTeamId := offer.OfferingTeam.Id

	dto := DtoFromOffer(offer)

	repo.g.Model(&OfferDto{}).
		Where("player_id = ? and offering_team_id = ?", playerId, offeringTeamId).
		Updates(dto)
}

func (repo *MarketRepo) DeleteOffer(o *models.Offer) {
	if o == nil {
		//TODO: this is for a double deleting offer, try to prevent this
		return
	}
	trx := repo.g.Where("player_id = ? and offering_team_id = ?", o.Player.Id, o.OfferingTeam.Id).Delete(&OfferDto{})
	if trx.Error != nil {
		panic("ERROR")
	}
}

func (repo *MarketRepo) GetOffersByPlayerTeamId(playerId, offeringTeamId string) (*models.Offer, bool) {
	var offer OfferDto

	trx := repo.g.Model(&OfferDto{}).Where("player_id = ? and offering_team_id = ?", playerId, offeringTeamId).
		Preload(teamRel).
		Preload(playerRel).
		Preload("OfferingTeam").
		Find(&offer)
	if trx.RowsAffected != 1 {
		return nil, false
	}

	return offer.Offer(), true
}

func (repo *MarketRepo) GetOffersByOfferingTeamId(offeringTeamId string) []*models.Offer {
	var dtos []OfferDto

	repo.g.Model(&OfferDto{}).Where("offering_team_id = ?", offeringTeamId).
		Preload(teamRel).
		Preload(playerRel).
		Preload("OfferingTeam").
		Find(&dtos)

	offers := make([]*models.Offer, len(dtos))
	for i, o := range dtos {
		offers[i] = o.Offer()
	}

	return offers
}

func (repo *MarketRepo) ApplyTransfer(o *models.Offer) {
	var player PlayerDto
	var offeringTeam TeamDto
	repo.g.Model(&PlayerDto{}).Preload(teamRel).Find(&player, o.Player.Id)
	repo.g.Model(&TeamDto{}).Find(&offeringTeam, o.OfferingTeam.Id)

	var playerPreviousTeam *TeamDto
	if player.Team != nil {
		playerPreviousTeam = player.Team
	}

	player.TeamId = &o.OfferingTeam.Id
	player.Wage = o.WageValue.Val
	player.YContract = *o.YContract

	if playerPreviousTeam != nil {
		playerPreviousTeam.Balance += o.BidValue.Value()
		offeringTeam.Balance -= o.BidValue.Value()
	}

	//TODO: check if balances are ok now

	//TODO: check if the transfer is between seasons or after seasons
	// in case it is inbetween drop half season log in history
	// o.TransferDate

}
