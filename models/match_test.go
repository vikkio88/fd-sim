package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchBuilder(t *testing.T) {
	rng := libs.NewRng(0)

	g := generators.NewTeamGenSeeded(rng)
	home := g.Team(enums.IT)
	away := g.Team(enums.EN)

	m := models.NewMatch(home, away)
	assert.Equal(t, home.Id, m.Home.Id)
	assert.Equal(t, away.Id, m.Away.Id)
	assert.NotNil(t, m.LineupHome)
	assert.NotNil(t, m.LineupAway)
}
func TestMatchResultIfPlayed(t *testing.T) {
	rng := libs.NewRng(0)

	g := generators.NewTeamGenSeeded(rng)
	home := g.Team(enums.IT)
	away := g.Team(enums.EN)

	m := models.NewMatch(home, away)
	_, ok := m.Result()
	assert.False(t, ok)

	m.Simulate(rng)

	r, ok := m.Result()
	assert.True(t, ok)
	assert.IsType(t, models.Result{}, *r)

	m.Simulate(rng)
	r2, _ := m.Result()
	assert.Equal(t, r, r2)
}

// func TestMultipleMatches(t *testing.T) {
// 	rng := libs.NewRng(0)

// 	g := generators.NewTeamGenSeeded(rng)
// 	home := g.Team(enums.IT)
// 	away := g.Team(enums.EN)

// 	for i := 0; i < 10000; i++ {
// 		m := models.NewMatch(home, away)
// 		m.Simulate(rng)
// 		res, _ := m.Result()

// 		fmt.Printf("%s - %s  %d - %d\n", m.Home.Name, m.Away.Name, res.GoalsHome, res.GoalsAway)
// 	}

// }
