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

	// calculating maluses
	goalsH, aMalus := m.LineupAway.Malus()
	goalsA, hMalus := m.LineupHome.Malus()

	goalsH -= hMalus
	goalsA -= aMalus

	startChance, homeStart := diffChance(m.LineupHome.teamStats.Skill, m.LineupAway.teamStats.Skill)
	if rng.ChanceI(startChance) {
		goalsH += rng.UInt(0, homeStart)
	}

	startChance, awayStart := diffChance(m.LineupAway.teamStats.Skill, m.LineupHome.teamStats.Skill)
	if rng.ChanceI(startChance) {
		goalsA += rng.UInt(0, awayStart)
	}

	if rng.ChanceI(int(m.LineupHome.teamStats.Morale)) {
		goalsH += 1
	}

	if rng.ChanceI(int(m.LineupAway.teamStats.Morale)) {
		goalsA += 1
	}

	goalsH, goalsA = m.strikersBonus(rng, goalsH, goalsA)

	goalsA, goalsH = m.defendersBonus(rng, goalsA, goalsH)

	goalsA, goalsH = m.normaliseGoals(goalsA, goalsH)

	scorersH := []string{}
	scorersA := []string{}
	m.result = NewResult(goalsH, goalsA, scorersH, scorersA)
}

func (*Match) normaliseGoals(goalsA int, goalsH int) (int, int) {
	if goalsA < 0 {
		goalsA = 0
	}

	if goalsH < 0 {
		goalsH = 0
	}
	return goalsA, goalsH
}

func (m *Match) defendersBonus(rng *libs.Rng, goalsA int, goalsH int) (int, int) {
	if bestDf, ok := m.LineupHome.BestPlayerInRole(DF); ok {
		if rng.ChanceI(bestDf.Skill) {
			goalsA -= rng.UInt(0, goalsA)
		}
	}

	if bestDf, ok := m.LineupAway.BestPlayerInRole(DF); ok {
		if rng.ChanceI(bestDf.Skill) {
			goalsH -= rng.UInt(0, goalsH)
		}
	}
	return goalsA, goalsH
}

func (m *Match) strikersBonus(rng *libs.Rng, goalsH int, goalsA int) (int, int) {
	if bestSt, ok := m.LineupHome.BestPlayerInRole(ST); ok {
		if rng.ChanceI(bestSt.Skill) {
			goalsH += rng.UInt(0, 2)
		}
	}

	if bestSt, ok := m.LineupAway.BestPlayerInRole(ST); ok {
		if rng.ChanceI(bestSt.Skill) {
			goalsA += rng.UInt(0, 2)
		}
	}
	return goalsH, goalsA
}

func (m *Match) Result() (*Result, bool) {
	if m.result == nil {
		return nil, false
	}

	return m.result, true
}

func (m *Match) String() string {
	res := ""
	if r, ok := m.Result(); ok {
		res = fmt.Sprintf("%d - %d", r.GoalsHome, r.GoalsAway)
	}

	return fmt.Sprintf("%s - %s %s", m.Home.Name, m.Away.Name, res)
}

const (
	chanceMin = 50
	chanceMax = 70
	goalRatio = 33.3
)

func diffChance(a, b float64) (int, int) {
	startingGoals := 0
	diff := a - b
	switch {
	case diff <= 10:
		startingGoals = 1
	case diff <= 40:
		startingGoals = 2
	case diff > 40:
		startingGoals = 3
	default:
		startingGoals = 1
	}

	return int(math.Min(100, math.Max(chanceMin, chanceMin+chanceMax*(b-a)/100))), startingGoals
}
