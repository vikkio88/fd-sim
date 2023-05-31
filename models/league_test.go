package models_test

import (
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeagueBuilder(t *testing.T) {
	//l := models.NewLeague()
	//assert.Equal(t, l, false)
}

func TestCalendarBuilder(t *testing.T) {
	teamIds := []string{"Juventus", "Milan"}
	calendar := models.NewRoundsCalendar(teamIds)
	assert.Len(t, calendar, 2)
	for i, r := range calendar {
		assert.Equal(t, i, r.Index)
		assert.Len(t, r.Matches, 1)
	}

	teamIds = []string{"Juventus", "Milan", "Crotone", "Palermo"}
	calendar = models.NewRoundsCalendar(teamIds)
	assert.Len(t, calendar, 6)
	for i, r := range calendar {
		assert.Equal(t, i, r.Index)
		assert.Len(t, r.Matches, 2)
	}
}
