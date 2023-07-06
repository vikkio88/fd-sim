package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRosterBuilder(t *testing.T) {
	r := models.NewRoster()
	p := models.NewPlayer("a", "a", 20, enums.DE, models.DF)
	p.SetVals(10, 10, 10)
	r.AddPlayer(&p)
	assert.Equal(t, 1, r.Len())
	assert.Equal(t, 10.0, r.AvgSkill())
	assert.Equal(t, 20.0, r.AvgAge())
	assert.Equal(t, "0.00 €", r.Wages().String())

	p2 := models.NewPlayer("b", "b", 30, enums.IT, models.DF)
	p2.SetVals(50, 10, 10)
	p2.Wage = utils.NewEuros(1000)
	// TODO: add here support to other currency
	// p2.Wage = utils.NewEuros(1000)
	r.AddPlayer(&p2)
	assert.Equal(t, 2, r.Len())
	assert.Equal(t, 30.0, r.AvgSkill())
	assert.Equal(t, 25.0, r.AvgAge())
	assert.Equal(t, "1000.00 €", r.Wages().String())

	assert.Equal(t, 2, len(r.InRole(models.DF)))
	assert.Equal(t, 2, len(r.IdsInRole(models.DF)))
	assert.Equal(t, 0, len(r.InRole(models.GK)))
	assert.Equal(t, 0, len(r.IdsInRole(models.GK)))

	assert.Greater(t, r.InRole(models.DF)[0].Skill.Val(), r.InRole(models.DF)[1].Skill.Val())

	p3 := models.NewPlayer("b", "b", 20, enums.IT, models.GK)
	p3.SetVals(60, 10, 10)
	r.AddPlayer(&p3)
	assert.Equal(t, 1, len(r.InRole(models.GK)))
	assert.Equal(t, 1, len(r.IdsInRole(models.GK)))
}

func TestLineupGeneration(t *testing.T) {
	tg := generators.NewTeamGen(0)

	team := tg.Team(enums.IT)

	lineup := team.Roster.Lineup(models.M442, 0, 100)
	count := countPlayersInLineup(lineup)
	assert.Equal(t, 11, count)
	lineup = team.Roster.Lineup(models.M343, 0, 100)
	count = countPlayersInLineup(lineup)
	assert.Equal(t, 11, count)
}

func TestLineupGetScorer(t *testing.T) {
	tg := generators.NewTeamGen(0)

	team := tg.Team(enums.IT)

	lineup := team.Roster.Lineup(models.M442, 0, 100)
	rng := libs.NewRng(0)
	c := models.NewEmptyRoleCounter()
	for i := 0; i < 1000; i++ {
		scorerId := lineup.Scorer(rng)
		p, _ := team.Roster.Player(scorerId)
		c[p.Role]++
	}
	assert.Equal(t, 0, c[models.GK])
	assert.Greater(t, c[models.ST], c[models.MF])
	assert.Greater(t, c[models.MF], c[models.DF])
}

func countPlayersInLineup(lineup *models.Lineup) int {
	count := 0
	for _, p := range lineup.Starting {
		count += len(p)
	}
	return count
}
