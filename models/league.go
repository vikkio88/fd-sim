package models

import (
	"fdsim/enums"
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
	Country enums.Country
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
	teamMap := TeamMap{}
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
func NewLeague(teams []*Team, seasonStartDate time.Time) *League {
	if len(teams)%2 != 0 {
		panic("Teams need to be an even number!")
	}
	teamMap, teamIds := TeamMapAndIdsFromTeams(teams)
	rounds := NewRoundsCalendar(teamIds, seasonStartDate.Year())
	return &League{
		Idable:      NewIdable(leagueIdGenerator()),
		Name:        "PLACEHOLDER",
		TeamMap:     teamMap,
		teams:       teams,
		Table:       NewTable(teams),
		Rounds:      rounds,
		totalRounds: (len(teams) * 2) - 2,
		RPointer:    0,
	}
}

func (l *League) UpdateLocales(name string, country enums.Country) {
	l.Name = name
	l.Country = country

}

func TeamMapAndIdsFromTeams(teams []*Team) (TeamMap, []string) {
	teamMap := TeamMap{}
	teamIds := make([]string, len(teams))
	for i, t := range teams {
		teamMap[t.Id] = t
		teamIds[i] = t.Id
	}
	return teamMap, teamIds
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

func (l *League) TableRow(index int) *TPHRow {
	team, row := l.Table.Get(index)
	return &TPHRow{
		Index: index,
		Team:  l.TeamMap[team].PH(),
		Row:   row,
	}
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
