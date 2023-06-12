package models

import (
	"fdsim/libs"
	"fmt"
	"math"

	"github.com/oklog/ulid/v2"
)

const (
	chanceMin  = 50
	chanceMax  = 70
	goalRatio  = 33.3
	matureTeam = 28.0
	youngTeam  = 22.0
	moraleHigh = 70.0
)

const matchInMemoryId = "mId"

func matchIdGenerator() string {
	return fmt.Sprintf("%s_%s", matchInMemoryId, ulid.Make())
}

type Match struct {
	Idable
	Home       TPH
	Away       TPH
	LineupHome *Lineup
	LineupAway *Lineup
	result     *Result
}

type MatchResult struct {
	Id         string
	Home       TPH
	Away       TPH
	LineupHome []string
	LineupAway []string
	Result     *Result
}

func NewMatchOnlyWithId(Id string) *Match {
	return &Match{
		Idable: NewIdable(Id),
	}
}

func NewMatchWithId(Id string, home, away *Team) *Match {
	return &Match{
		Idable:     NewIdable(Id),
		Home:       home.PH(),
		Away:       away.PH(),
		LineupHome: home.Lineup(),
		LineupAway: away.Lineup(),
	}
}
func NewMatch(home, away *Team) *Match {
	return NewMatchWithId(matchIdGenerator(), home, away)
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

	startChance, homeStart := startingGoals(m.LineupHome.teamStats.Skill, m.LineupAway.teamStats.Skill)
	if rng.ChanceI(startChance) {
		goalsH += rng.UInt(1, homeStart)
	}

	startChance, awayStart := startingGoals(m.LineupAway.teamStats.Skill, m.LineupHome.teamStats.Skill)
	if rng.ChanceI(startChance) {
		goalsA += rng.UInt(1, awayStart)
	}

	if rng.ChanceI(int(m.LineupHome.teamStats.Morale)) {
		goalsH += rng.UInt(0, 1)
		goalsA -= rng.UInt(0, 1)
	}

	if rng.ChanceI(int(m.LineupAway.teamStats.Morale)) {
		goalsA += rng.UInt(0, 1)
		goalsH -= rng.UInt(0, 1)
	}

	goalsH, goalsA = m.bestStrikerBonus(rng, goalsH, goalsA)
	goalsA, goalsH = m.bestDefenderBonus(rng, goalsA, goalsH)

	goalsA, goalsH = m.defenceBonus(rng, goalsA, goalsH)
	goalsA, goalsH = m.attackBonus(rng, goalsA, goalsH)

	goalsA, goalsH = m.lineupMoraleAgeBonus(rng, goalsA, goalsH)

	goalsH, goalsA = m.fluke(goalsH, rng, goalsA)

	goalsA, goalsH = m.normaliseGoals(goalsA, goalsH, rng)

	scorersH := m.LineupHome.Scorers(goalsH, rng)
	scorersA := m.LineupAway.Scorers(goalsA, rng)
	m.result = NewResult(goalsH, goalsA, scorersH, scorersA)
}

func (*Match) fluke(goalsH int, rng *libs.Rng, goalsA int) (int, int) {
	goalsH += rng.PlusMinusVal(1, 50)
	goalsA += rng.PlusMinusVal(1, 50)
	return goalsH, goalsA
}

func (m *Match) lineupMoraleAgeBonus(rng *libs.Rng, goalsA int, goalsH int) (int, int) {
	if m.LineupHome.lineupStats.Morale >= moraleHigh {
		goalsH += rng.UInt(0, 2)
	}

	if m.LineupAway.lineupStats.Morale >= moraleHigh {
		goalsA += rng.UInt(0, 2)
	}

	if rng.ChanceI(int(m.LineupHome.lineupStats.Morale)) {
		goalsA -= rng.UInt(0, 2)
	}

	if rng.ChanceI(int(m.LineupAway.lineupStats.Morale)) {
		goalsH -= rng.UInt(0, 2)
	}

	if m.LineupHome.lineupStats.Age >= matureTeam {
		goalsA -= rng.UInt(0, 1)
	}

	if m.LineupAway.lineupStats.Age >= matureTeam {
		goalsH -= rng.UInt(0, 1)
	}

	if m.LineupHome.lineupStats.Age <= youngTeam {
		goalsH += rng.UInt(0, 1)
	}

	if m.LineupAway.lineupStats.Age >= youngTeam {
		goalsA += rng.UInt(0, 1)
	}

	return goalsA, goalsH
}

func (m *Match) defenceBonus(rng *libs.Rng, goalsA int, goalsH int) (int, int) {
	if rng.ChanceI(diffChance(m.LineupHome.sectorStat[DF].Skill, m.LineupAway.sectorStat[ST].Skill)) {
		goalsA -= rng.UInt(1, 2)
	}

	if rng.ChanceI(diffChance(m.LineupAway.sectorStat[DF].Skill, m.LineupHome.sectorStat[ST].Skill)) {
		goalsH -= rng.UInt(1, 2)
	}
	return goalsA, goalsH
}

func (m *Match) attackBonus(rng *libs.Rng, goalsA int, goalsH int) (int, int) {
	if rng.ChanceI(diffChance(m.LineupHome.sectorStat[ST].Skill, m.LineupAway.sectorStat[DF].Skill)) {
		goalsH += rng.UInt(0, 1)
	}

	if rng.ChanceI(diffChance(m.LineupAway.sectorStat[ST].Skill, m.LineupHome.sectorStat[DF].Skill)) {
		goalsA += rng.UInt(0, 1)
	}

	return goalsA, goalsH
}

func (m *Match) normaliseGoals(goalsA, goalsH int, rng *libs.Rng) (int, int) {
	if goalsA < 0 {
		goalsA = 0
	}

	if goalsA > 7 && rng.ChanceI(60) {
		goalsA -= rng.UInt(1, 3)
	}

	if goalsH < 0 {
		goalsH = 0
	}

	if goalsH > 7 && rng.ChanceI(60) {
		goalsH -= rng.UInt(1, 3)
	}

	return goalsA, goalsH
}

func (m *Match) bestDefenderBonus(rng *libs.Rng, goalsA int, goalsH int) (int, int) {
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

func (m *Match) bestStrikerBonus(rng *libs.Rng, goalsH int, goalsA int) (int, int) {
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

func (m *Match) SetResult(result *Result) {
	if m.result != nil {
		return
	}

	m.result = result
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

func startingGoals(a, b float64) (int, int) {
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

	return diffChance(a, b), startingGoals
}

func diffChance(a, b float64) int {
	return int(math.Min(100, math.Max(chanceMin, chanceMin+chanceMax*(b-a)/100)))
}
