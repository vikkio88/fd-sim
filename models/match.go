package models

import (
	"fdsim/libs"
	"fmt"
	"math"

	"github.com/oklog/ulid/v2"
)

const matchInMemoryId = "mId"

func matchIdGenerator() string {
	return fmt.Sprintf("%s_%s", matchInMemoryId, ulid.Make())
}

type Match struct {
	idable
	Home       TPH
	Away       TPH
	LineupHome *Lineup
	LineupAway *Lineup
	result     *Result
}

func NewMatch(home, away *Team) *Match {
	return &Match{
		idable:     NewIdable(matchIdGenerator()),
		Home:       home.PH(),
		Away:       away.PH(),
		LineupHome: home.Lineup(),
		LineupAway: away.Lineup(),
	}
}

func (m *Match) Simulate(rng *libs.Rng) {
	if m.result != nil {
		return
	}

	goalsH, goalsA := 0, 0
	if rng.ChanceI(diffChance(m.LineupHome.teamStats.Skill, m.LineupAway.teamStats.Skill)) {
		goalsH += 1
	}

	if rng.ChanceI(diffChance(m.LineupAway.teamStats.Skill, m.LineupHome.teamStats.Skill)) {
		goalsA += 1
	}

	if rng.ChanceI(int(m.LineupHome.teamStats.Morale)) {
		goalsH += 1
	}

	if rng.ChanceI(int(m.LineupAway.teamStats.Morale)) {
		goalsA += 1
	}

	scorersH := []string{}
	scorersA := []string{}
	m.result = NewResult(goalsH, goalsA, scorersH, scorersA)
}

func (m *Match) Result() (*Result, bool) {
	if m.result == nil {
		return nil, false
	}

	return m.result, true
}

const (
	chanceMin = 40
	chanceMax = 60
)

func diffChance(a, b float64) int {
	return int(math.Min(100, math.Max(chanceMin, chanceMin+chanceMax*(b-a)/100)))
}
