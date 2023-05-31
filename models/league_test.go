package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
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

func TestLeagueSimulation(t *testing.T) {
	rng := libs.NewRng(5)
	tg := generators.NewTeamGenSeeded(rng)

	ts := tg.TeamsWithCountry(4, enums.IT)

	l := models.NewLeague("Serie A", ts)

	r, ok := l.NextRound()
	assert.True(t, ok)
	r.Simulate(rng)
	l.Update(r)

	r, ok = l.NextRound()
	assert.True(t, ok)
	r.Simulate(rng)
	l.Update(r)

	r, ok = l.NextRound()
	assert.True(t, ok)
	r.Simulate(rng)
	l.Update(r)

	r, ok = l.NextRound()
	assert.True(t, ok)
	r.Simulate(rng)
	l.Update(r)

	r, ok = l.NextRound()
	assert.True(t, ok)
	r.Simulate(rng)
	l.Update(r)

	r, ok = l.NextRound()
	assert.True(t, ok)
	r.Simulate(rng)
	l.Update(r)

	r, ok = l.NextRound()
	assert.False(t, ok)
	assert.Nil(t, r)

	// for _, r := range l.Table.Rows() {
	// 	fmt.Printf("%s\n", r)
	// }
}
