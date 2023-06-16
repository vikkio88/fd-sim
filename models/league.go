package models

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

// Round Placeholder
type RPH struct {
	Id      string
	Index   int
	Matches []MPH
}

func NewRoundsCalendar(teamIds []string) []RPH {
	var firstHalf []RPH
	var secondHalf []RPH
	numberOfTeams := len(teamIds)
	totalRounds := numberOfTeams - 1
	matchesPerRound := numberOfTeams / 2

	for round := 0; round < totalRounds; round++ {
		var tempRoundFirstHalf []MPH
		var tempRoundSecondHalf []MPH

		for match := 0; match < matchesPerRound; match++ {
			home := (round + match) % (numberOfTeams - 1)
			away := (numberOfTeams - 1 - match + round) % (numberOfTeams - 1)

			if match == 0 {
				away = numberOfTeams - 1
			}

			tempRoundFirstHalf = append(tempRoundFirstHalf, MPH{
				Id:   matchIdGenerator(),
				Home: teamIds[home],
				Away: teamIds[away],
			})

			tempRoundSecondHalf = append(tempRoundSecondHalf, MPH{
				Id:   matchIdGenerator(),
				Home: teamIds[away],
				Away: teamIds[home],
			})
		}

		firstHalf = append(firstHalf, RPH{
			Id:      roundIdGenerator(),
			Index:   round,
			Matches: tempRoundFirstHalf,
		})

		secondHalf = append(secondHalf, RPH{
			Id:      roundIdGenerator(),
			Index:   round + totalRounds,
			Matches: tempRoundSecondHalf,
		})
	}

	firstHalf = append(firstHalf, secondHalf...)

	return firstHalf
}

func (r *RPH) Round(teamsMap map[string]*Team) *Round {
	matches := make([]*Match, len(r.Matches))
	for i, mph := range r.Matches {
		home, ok := teamsMap[mph.Home]
		away, ok2 := teamsMap[mph.Away]
		if !ok || !ok2 {
			panic("Empty team map")
		}
		matches[i] = mph.Match(mph.Id, home, away)

	}
	return NewRound(r.Id, r.Index, matches)
}

func (r *RPH) RoundTPH(teamsMap map[string]*Team) *RPHTPH {
	matches := make([]MPHTPH, len(r.Matches))
	for i, mph := range r.Matches {
		home, ok := teamsMap[mph.Home]
		away, ok2 := teamsMap[mph.Away]
		if !ok || !ok2 {
			panic("Empty team map")
		}
		matches[i] = *mph.MPHTPH(mph.Id, home.PH(), away.PH())

	}
	return &RPHTPH{Id: r.Id, Index: r.Index, Matches: matches}
}

// Match Placeholder
type MPH struct {
	Id   string
	Home string
	Away string
}

// Round PH With Team PH
type RPHTPH struct {
	Id      string
	Index   int
	Matches []MPHTPH
}

// Match with TeamPH
type MPHTPH struct {
	Id     string
	Home   TPH
	Away   TPH
	Result *Result
}

func (m *MPH) MPHTPH(Id string, home, away TPH) *MPHTPH {
	return &MPHTPH{
		Id:   m.Id,
		Home: home,
		Away: away,
	}
}

func (r *MPH) Match(Id string, home, away *Team) *Match {
	return NewMatchWithId(Id, home, away)
}

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

func NewLeague(name string, teams []*Team) League {
	if len(teams)%2 != 0 {
		panic("Teams need to be an even number!")
	}
	teamMap := map[string]*Team{}
	teamIds := make([]string, len(teams))
	for i, t := range teams {
		teamMap[t.Id] = t
		teamIds[i] = t.Id
	}
	rounds := NewRoundsCalendar(teamIds)
	return League{
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

func (l *League) NextRound() (*Round, bool) {
	if l.IsFinished() {
		return nil, false
	}
	return l.Rounds[l.RPointer].Round(l.TeamMap), true
}

func (l *League) Update(round *Round) {
	l.Table.Update(round)
	//l. Update Stats
	l.RPointer++
}
