package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineupBuilder(t *testing.T) {
	l := models.NewLineup(models.M442, models.NewRolePPHMap(), models.TeamStats{})
	assert.IsType(t, models.Lineup{}, *l)
}

func TestLineupPlayers(t *testing.T) {
	g := generators.NewTeamGen(0)

	team := g.Team(enums.IT)
	l := team.Lineup()

	bestDF, ok := l.BestPlayerInRole(models.DF)
	assert.True(t, ok)
	assert.GreaterOrEqual(t, bestDF.Skill, l.Starting[models.DF][1].Skill)

}
