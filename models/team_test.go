package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamBuilder(t *testing.T) {
	team := models.NewTeam("Juventus", "Torino", enums.IT)
	assert.Equal(t, "Juventus", team.Name)
	assert.Equal(t, "Torino", team.City)
	assert.Equal(t, "Italian", team.Country.Nationality())

	assert.IsType(t, utils.Money{}, team.Balance)
	assert.Equal(t, 0.0, team.TransferRatio)
	assert.Nil(t, team.Coach)
	assert.NotNil(t, team.Roster)
}

func TestTeamAcceptingOffer(t *testing.T) {
	tg := generators.NewTeamGen(0)
	team := tg.Team(enums.EN)

	p := team.Roster.Players()[0]
	pVal := p.Value.Value()
	p.YContract = 5
	p.Age = 20

	acceptedPerc := team.OfferAcceptanceChance(utils.NewEurosFromF(pVal-(pVal/2)), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 50)

	p.YContract = 1
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal-(pVal/2)), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 60)
	assert.GreaterOrEqual(t, acceptedPerc.Val(), 50)

	p.YContract = 5
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 90)

	p.YContract = 1
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 100)

	p.YContract = 5
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal+(pVal/2)), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 90)

	p.YContract = 1
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal+(pVal/2)), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 100)

	//making player older
	p.Age = 30
	p.YContract = 5
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal-(pVal/2)), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 50)

	p.YContract = 1
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal-(pVal/2)), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 75)
	assert.GreaterOrEqual(t, acceptedPerc.Val(), 50)

	p.YContract = 5
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 100)

	p.YContract = 1
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 100)

	p.YContract = 5
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal+(pVal/2)), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 100)

	p.YContract = 1
	acceptedPerc = team.OfferAcceptanceChance(utils.NewEurosFromF(pVal+(pVal/2)), p.Id)
	assert.LessOrEqual(t, acceptedPerc.Val(), 100)
}
