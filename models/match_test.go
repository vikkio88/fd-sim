package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"fmt"
	"sort"
	"testing"
	"time"

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

func TestMatchSimulationWithMissingPlayers(t *testing.T) {
	rng := libs.NewRng(time.Now().Unix())

	g := generators.NewTeamGenSeeded(rng)
	gp := generators.NewPeopleGenSeeded(rng)
	home := g.Team(enums.IT)
	away := models.NewTeam("Broken", "", enums.IT)
	away.Roster.AddPlayer(gp.PlayerWithRole(enums.IT, models.GK))

	m := models.NewMatch(home, away)
	m.Simulate(rng)

	fmt.Println(m)
}

func TestMultipleMatches(t *testing.T) {
	// t.Skip("Long Test")
	rng := libs.NewRng(100)

	g := generators.NewTeamGenSeeded(rng)
	home := g.Team(enums.IT)
	away := g.Team(enums.EN)
	for _, r := range []models.Role{models.GK, models.GK, models.DF, models.DF, models.DF, models.MF, models.MF, models.ST, models.ST, models.ST} {
		crazyP := models.NewPlayer("a", "a", 10, enums.IT, r)
		crazyP.SetVals(10, 10, 100)
		home.Roster.AddPlayer(&crazyP)
	}

	gr := map[string]int{}
	points := map[string]int{}
	won := map[string]int{}
	matches := 10000

	for i := 0; i < matches; i++ {
		m := models.NewMatch(home, away)
		m.Simulate(rng)
		res, _ := m.Result()
		scoreStr := fmt.Sprintf("%d - %d", res.GoalsHome, res.GoalsAway)
		gr[scoreStr]++
		fmt.Printf("%s\n", m)
		if res.Draw() {
			points[home.Name] += 1
			points[away.Name] += 1
		} else if res.HomeWin() {
			points[home.Name] += 3
			won[home.Name]++
		} else {
			points[away.Name] += 3
			won[away.Name]++
		}
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

	fmt.Printf(
		"%s %.2f: %d => ppm: %.2f , w%% %.2f\n",
		home, home.Roster.AvgSkill(),
		points[home.Name],
		float32(points[home.Name])/float32(matches),
		float32(won[home.Name])/float32(matches),
	)
	fmt.Printf(
		"%s %.2f: %d => ppm: %.2f , w%% %.2f\n",
		away, away.Roster.AvgSkill(),
		points[away.Name],
		float32(points[away.Name])/float32(matches),
		float32(won[away.Name])/float32(matches),
	)

	fmt.Printf("drawn perc: %.2f\n", 1.0-(float32(won[away.Name])/float32(matches)+float32(won[home.Name])/float32(matches)))
}
