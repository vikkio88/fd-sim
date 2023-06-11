package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestLeagueBuilder(t *testing.T) {
	rng := libs.NewRng(5)
	tg := generators.NewTeamGenSeeded(rng)
	ts := tg.TeamsWithCountry(4, enums.IT)
	l := models.NewLeague("Serie A", ts)
	assert.Equal(t, len(l.Rounds), ((4 - 1) * 2))

	ts = tg.TeamsWithCountry(3, enums.IT)
	assertPanic(t, func() {
		models.NewLeague("Serie A", ts)
	})
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
	teams := 4
	rng := libs.NewRng(time.Now().Unix())
	tg := generators.NewTeamGenSeeded(rng)
	ts := tg.TeamsWithCountry(teams, enums.IT)
	l := models.NewLeague("Serie A", ts)

	for i := 0; i < len(l.Rounds); i++ {
		assert.False(t, l.IsFinished())
		r, ok := l.NextRound()
		assert.True(t, ok)
		r.Simulate(rng)
		printResults(r)
		l.Update(r)
	}

	assert.True(t, l.IsFinished())
	r, ok := l.NextRound()
	assert.False(t, ok)
	assert.Nil(t, r)

	println("\nTABLE")
	for i, r := range l.Table.Rows() {
		fmt.Printf("%d  - %s\n", i+1, r)
	}
}

func printResults(r *models.Round) {
	fmt.Printf("\nRound %d\n", r.Index+1)
	for _, m := range r.Matches {
		fmt.Printf("%s\n", m)
	}
}
