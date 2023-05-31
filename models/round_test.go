package models_test

import (
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundBuilder(t *testing.T) {
	rng := libs.NewRng(0)
	tg := generators.NewTeamGenSeeded(rng)
	teams := tg.Teams(4)
	matches := []*models.Match{
		models.NewMatch(teams[0], teams[1]),
		models.NewMatch(teams[2], teams[3]),
	}
	round := models.NewRound(0, matches)
	r, ok := round.Results()
	assert.Nil(t, r)
	assert.False(t, ok)

	round.Simulate(rng)
	r, ok = round.Results()
	assert.NotNil(t, r)
	assert.True(t, ok)
}
