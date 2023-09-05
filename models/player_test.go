package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildingPlayer(t *testing.T) {
	p := models.NewPlayer("Name", "Surname", 30, enums.EN, models.GK)
	assert.Equal(t, "Name", p.Name)
	assert.Equal(t, "Surname", p.Surname)
	assert.Equal(t, "English", p.Country.Nationality())
	assert.Equal(t, 30, p.Age)
	assert.Equal(t, "Goalkeeper", p.Role.String())

	// idable
	assert.Contains(t, p.Id, "pmId_")

	// Wage and Value
	assert.Equal(t, p.IdealWage.Val, utils.Money{}.Val)
	assert.Equal(t, p.Value.Val, utils.Money{}.Val)
}

func TestPlayerIsSkillable(t *testing.T) {
	p := models.NewPlayer("Mario", "Rossi", 17, enums.DE, models.ST)

	assert.IsType(t, utils.Perc{}, p.Skill)
	assert.IsType(t, utils.Perc{}, p.Morale)

	p.SetVals(50, 99, 10)

	assert.Equal(t, 50, p.Skill.Val())
	assert.Equal(t, 99, p.Morale.Val())
	assert.Equal(t, 10, p.Fame.Val())
}

func TestPlayerConctractDecision(t *testing.T) {
	pd := buildPlayerDetailed()
	td := buildTeamDetailed()

	money := utils.NewEuros(1)
	yContract := 1

	chance := pd.WageAcceptanceChance(money, yContract, td)
	assert.LessOrEqual(t, 5, chance.Val())

	// if matching ideal wage and morale is low
	money = pd.IdealWage
	yContract = 1
	pd.Morale = utils.NewPerc(0)
	chance = pd.WageAcceptanceChance(money, yContract, td)
	assert.LessOrEqual(t, 95, chance.Val())

	// if wage higher than current
	money = pd.Wage
	money.Modify(.1)
	yContract = 1
	pd.Morale = utils.NewPerc(60)
	chance = pd.WageAcceptanceChance(money, yContract, td)
	assert.GreaterOrEqual(t, 90, chance.Val())

	// if team is better but same wage
	money = pd.Wage
	gtd := buildGreatTeam()
	pd.Morale.SetVal(60)
	yContract = 1
	chance = pd.WageAcceptanceChance(money, yContract, gtd)
	assert.GreaterOrEqual(t, chance.Val(), 60)

}

func buildGreatTeam() *models.TeamDetailed {
	tg := generators.NewTeamGen(0)
	team := tg.Team(enums.EN)
	pg := generators.NewPeopleGen(0)

	ps := pg.Players(20)

	for _, p := range ps {
		p.Skill = utils.NewPerc(100)
	}
	td := &models.TeamDetailed{Team: *team}
	return td
}

func buildTeamDetailed() *models.TeamDetailed {
	tg := generators.NewTeamGen(0)
	team := tg.Team(enums.EN)
	td := &models.TeamDetailed{Team: *team}
	return td
}

func buildPlayerDetailed() models.PlayerDetailed {
	g := generators.NewPeopleGen(0)
	p := g.Player(enums.EN)
	pd := models.PlayerDetailed{Player: *p}
	return pd
}
