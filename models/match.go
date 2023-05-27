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
	Home TPH
	Away TPH
}

func NewMatch(home, away *Team) *Match {
	return &Match{
		idable: NewIdable(matchIdGenerator()),
		Home:   home.PH(),
		Away:   away.PH(),
	}
}
