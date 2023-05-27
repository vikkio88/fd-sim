package models

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

const matchInMemoryId = "mId"

func matchIdGenerator() string {
	return fmt.Sprintf("%s_%s", matchInMemoryId, ulid.Make())
}

type Match struct {
	idable
	Home   TPH
	Away   TPH
	result *Result
}

func NewMatch(home, away *Team) *Match {
	return &Match{
		idable: NewIdable(matchIdGenerator()),
		Home:   home.PH(),
		Away:   away.PH(),
	}
}

func (m *Match) Simulate() {
	m.result = NewResult(0, 0, []string{}, []string{})
}

func (m *Match) Result() (*Result, bool) {
	if m.result == nil {
		return nil, false
	}

	return m.result, true
}
