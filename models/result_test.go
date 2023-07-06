package models_test

import (
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeScorers []string

func TestResultBuilder(t *testing.T) {
	fakeHomeScore := models.PlayerScoreMap{
		"id1": 5.0,
		"id2": 10.0,
	}
	fakeAwayScore := models.PlayerScoreMap{
		"id11": 5.0,
		"id22": 5.0,
	}

	r := models.NewResult(1, 1, fakeScorers, fakeScorers, fakeHomeScore, fakeAwayScore)
	assert.True(t, r.Draw())
	assert.False(t, r.HomeWin())
	assert.False(t, r.AwayWin())
	assert.Equal(t, models.RX, r.X12())
	assert.Equal(t, 5.0, r.ScoreHome["id1"])
	assert.Equal(t, 10.0, r.ScoreHome["id2"])
	assert.Equal(t, 5.0, r.ScoreAway["id11"])
	assert.Equal(t, 5.0, r.ScoreAway["id22"])

	fakeHomeScore2 := models.PlayerScoreMap{
		"id1": 10.0,
		"id2": 10.0,
	}
	fakeAwayScore2 := models.PlayerScoreMap{
		"id11": 10.0,
		"id22": 10.0,
	}
	r = models.NewResult(1, 3, fakeScorers, fakeScorers, fakeHomeScore2, fakeAwayScore2)
	assert.False(t, r.Draw())
	assert.False(t, r.HomeWin())
	assert.True(t, r.AwayWin())
	assert.Equal(t, models.R2, r.X12())

	assert.Equal(t, 10.0, r.ScoreHome["id1"])
	assert.Equal(t, 10.0, r.ScoreHome["id2"])
	assert.Equal(t, 10.0, r.ScoreAway["id11"])
	assert.Equal(t, 10.0, r.ScoreAway["id22"])

	r = models.NewResult(3, 0, fakeScorers, fakeScorers, models.PlayerScoreMap{}, models.PlayerScoreMap{})
	assert.False(t, r.Draw())
	assert.True(t, r.HomeWin())
	assert.False(t, r.AwayWin())
	assert.Equal(t, models.R1, r.X12())
}
