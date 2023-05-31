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
	assert.Equal(t, r.GoalsHome, len(r.ScorersHome))
	assert.Equal(t, r.GoalsAway, len(r.ScorersAway))
	for _, id := range r.ScorersHome {
		p, ok := home.Roster.Player(id)
		assert.NotNil(t, p)
		assert.True(t, ok)
	}

	for _, id := range r.ScorersAway {
		p, ok := away.Roster.Player(id)
		assert.NotNil(t, p)
		assert.True(t, ok)
	}

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

	r, ok := m.Result()
	assert.NotNil(t, r)
	assert.True(t, ok)

	assert.True(t, r.HomeWin())
	assert.Greater(t, r.GoalsHome, 1)
	assert.Equal(t, 0, r.GoalsAway)
}

func TestMultipleMatches(t *testing.T) {
	t.Skip("Long Test")
	seeded := true

	var seed int64 = 100
	if !seeded {
		seed = time.Now().Unix()
	}

	rng := libs.NewRng(seed)

	g := generators.NewTeamGenSeeded(rng)
	home := g.Team(enums.IT)
	away := g.Team(enums.EN)
	for _, r := range []models.Role{models.GK, models.GK, models.DF, models.DF, models.DF, models.MF, models.MF, models.ST, models.ST, models.ST} {
		player := models.NewPlayer("a", "a", 10, enums.IT, r)
		player.SetVals(1, 10, 100)
		home.Roster.AddPlayer(&player)
	}

	for _, r := range []models.Role{models.GK, models.GK, models.DF, models.DF, models.DF, models.MF, models.MF, models.ST, models.ST, models.ST} {
		player := models.NewPlayer("a", "a", 25, enums.IT, r)
		player.SetVals(100, 100, 100)
		away.Roster.AddPlayer(&player)
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

	printTeam(home, points, won, matches)
	printTeam(away, points, won, matches)
	fmt.Printf("drawn perc: %.2f\n", 1.0-(float32(won[away.Name])/float32(matches)+float32(won[home.Name])/float32(matches)))
}

func printTeam(t *models.Team, points, won map[string]int, matchesCount int) {
	fmt.Printf(
		"%s s:%.2f m:%.2f a: %.2f: %d => ppm: %.2f , w%% %.2f\n",
		t,
		t.Roster.AvgSkill(),
		t.Roster.AvgMorale(),
		t.Roster.AvgAge(),
		points[t.Name],
		float32(points[t.Name])/float32(matchesCount),
		float32(won[t.Name])/float32(matchesCount),
	)
}
