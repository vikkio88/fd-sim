package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
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

func TestMultipleMatches(t *testing.T) {
	t.Skip("Long Test")
	rng := libs.NewRng(0)

	g := generators.NewTeamGenSeeded(rng)
	home := g.Team(enums.IT)
	away := g.Team(enums.EN)

	crazyP := models.NewPlayer("a", "a", 10, enums.IT, models.ST)
	crazyP.SetVals(100, 100, 100)

	home.Roster.AddPlayer(&crazyP)

	gr := map[string]int{}

	for i := 0; i < 10000; i++ {
		m := models.NewMatch(home, away)
		m.Simulate(rng)
		res, _ := m.Result()
		scoreStr := fmt.Sprintf("%d - %d", res.GoalsHome, res.GoalsAway)
		gr[scoreStr]++
		fmt.Printf("%s - %s  %s\n", m.Home.Name, m.Away.Name, scoreStr)
	}
	grK := maps.Keys(gr)

	sort.SliceStable(grK, func(i, j int) bool {
		return gr[grK[i]] > gr[grK[j]]
	})

	fmt.Println()
	for _, s := range grK {
		c := gr[s]

		fmt.Printf("%s : %d\n", s, c)
	}

}
