package models_test

import (
	"fdsim/conf"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fmt"
	"testing"
	"time"

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
	round := models.NewRound("", 0, time.Now(), matches)
	r, ok := round.Results()
	assert.Nil(t, r)
	assert.False(t, ok)

	round.Simulate(rng)
	r, ok = round.Results()
	assert.NotNil(t, r)
	assert.True(t, ok)
}

func TestRoundStats(t *testing.T) {
	rng := libs.NewRng(1230)
	tg := generators.NewTeamGenSeeded(rng)
	teams := tg.Teams(2)
	matches := []*models.Match{
		models.NewMatch(teams[0], teams[1]),
	}
	round := models.NewRound("", 0, time.Now(), matches)
	round.Simulate(rng)

	rows := models.StatsFromRoundResult(round, "leagueId")
	assert.Equal(t, 22, len(rows))
	// this should be 0-2
	res, _ := round.Matches[0].Result()
	goals := res.GoalsHome + res.GoalsAway
	//TODO: this sometimes goes to 0-0
	assert.GreaterOrEqual(t, goals, 1)
	goalsAcc := 0
	for _, r := range rows {
		goalsAcc += r.Goals
		assert.Equal(t, r.Played, 1)
	}
	assert.Equal(t, goals, goalsAcc)

	rng = libs.NewRng(0)
	matches = []*models.Match{
		models.NewMatch(teams[1], teams[0]),
	}
	round = models.NewRound("", 0, time.Now(), matches)
	round.Simulate(rng)
	res, _ = round.Matches[0].Result()
	goals2 := res.GoalsHome + res.GoalsAway
	assert.GreaterOrEqual(t, goals2, 1)
	rows2 := models.StatsFromRoundResult(round, "leagueId")
	goalsAcc2 := 0
	updatedRows := models.MergeStats(rows, rows2)
	for _, r := range updatedRows {
		goalsAcc2 += r.Goals
		assert.Equal(t, r.Played, 2)
	}

	assert.Equal(t, goalsAcc2, goals+goals2)
}

func TestRoundDates(t *testing.T) {
	t.Skip("Slow")
	teams := 18
	for i := 0; i <= teams-1; i++ {
		fmt.Printf(
			"Round %d - %s\n",
			i,
			models.
				GetRoundDateByIndex(2023, time.September, i, false).
				Format(conf.DateFormatGame),
		)

		fmt.Printf(
			"Round %d - %s\n",
			teams-1+i,
			models.
				GetRoundDateByIndex(2023, time.September, i, true).
				Format(conf.DateFormatGame),
		)
	}
}
