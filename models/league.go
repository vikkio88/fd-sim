package models

import (
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

type TeamMap map[string]*Team

type League struct {
	Idable
	Name    string
	TeamMap TeamMap
	teams   []*Team
	// Rounds Placeholders
	Rounds []RPH
	Table  *Table
	// Last round pointer
	RPointer    int
	totalRounds int
}

const leagueInMemoryId = "leId"

func leagueIdGenerator() string {
	return fmt.Sprintf("%s_%s", leagueInMemoryId, ulid.Make())
}

func NewLeagueWithData(id, name string, teams []*Team) *League {
	teamMap := map[string]*Team{}
	teamIds := make([]string, len(teams))
	for i, t := range teams {
		teamMap[t.Id] = t
		teamIds[i] = t.Id
	}
	return &League{
		Idable:      NewIdable(id),
		Name:        name,
		totalRounds: (len(teams) * 2) - 2,
		TeamMap:     teamMap,
		teams:       teams,
	}
}

// SeasonStart comes from League Generation at the beginning of the League
func NewLeague(name string, teams []*Team, seasonStartDate time.Time) *League {
	if len(teams)%2 != 0 {
		panic("Teams need to be an even number!")
	}
	teamMap := map[string]*Team{}
	teamIds := make([]string, len(teams))
	for i, t := range teams {
		teamMap[t.Id] = t
		teamIds[i] = t.Id
	}
	rounds := NewRoundsCalendar(teamIds, seasonStartDate.Year())
	return &League{
		Idable:      NewIdable(leagueIdGenerator()),
		Name:        name,
		TeamMap:     teamMap,
		teams:       teams,
		Table:       NewTable(teams),
		Rounds:      rounds,
		totalRounds: (len(teams) * 2) - 2,
		RPointer:    0,
	}
}

func (l *League) RoundsPH() []*RPHTPH {
	rounds := make([]*RPHTPH, len(l.Rounds))
	for i, r := range l.Rounds {
		rounds[i] = r.RoundTPH(l.TeamMap)
	}
	return rounds
}

func (l *League) TableRows() []*TPHRow {
	return l.Table.TPHRows(l.TeamMap)
}

func (l *League) Teams() []*Team {
	return l.teams
}

func (l *League) IsFinished() bool {
	return l.RPointer >= l.totalRounds
}

// returns the next round and a bool representing whether there are more rounds
func (l *League) NextRound() (*Round, bool) {
	if l.IsFinished() {
		return nil, false
	}
	return l.Rounds[l.RPointer].Round(l.TeamMap), true
}

func (l *League) Update(round *Round) {
	l.Table.Update(round)
	l.RPointer++
}
