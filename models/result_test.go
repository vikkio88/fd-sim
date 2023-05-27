package models_test

import (
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeScorers []string

func TestResultBuilder(t *testing.T) {
	r := models.NewResult(1, 1, fakeScorers, fakeScorers)
	assert.True(t, r.Draw())
	assert.False(t, r.HomeWin())
	assert.False(t, r.AwayWin())
	assert.Equal(t, models.RX, r.X12())

	r = models.NewResult(1, 3, fakeScorers, fakeScorers)
	assert.False(t, r.Draw())
	assert.False(t, r.HomeWin())
	assert.True(t, r.AwayWin())
	assert.Equal(t, models.R2, r.X12())

	r = models.NewResult(3, 0, fakeScorers, fakeScorers)
	assert.False(t, r.Draw())
	assert.True(t, r.HomeWin())
	assert.False(t, r.AwayWin())
	assert.Equal(t, models.R1, r.X12())
}
