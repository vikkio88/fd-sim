package models

import (
	"fdsim/libs"
	"fdsim/utils"
	"fmt"

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
	if rng.Chance(utils.NewPerc(int(m.LineupHome.teamStats.Skill - m.LineupAway.teamStats.Skill))) {
		goalsH += 1
	}

	if rng.Chance(utils.NewPerc(int(m.LineupAway.teamStats.Skill - m.LineupAway.teamStats.Skill))) {
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
