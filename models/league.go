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
				Home: teamIds[home],
				Away: teamIds[away],
			})

			tempRoundSecondHalf = append(tempRoundSecondHalf, MPH{
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
		home := teamsMap[mph.Home]
		away := teamsMap[mph.Away]
		matches[i] = mph.Match(mph.Id, home, away)

	}
	return NewRound(r.Id, r.Index, matches)
}

// Match Placeholder
type MPH struct {
	Id   string
	Home string
	Away string
}

func (r *MPH) Match(Id string, home, away *Team) *Match {
	return NewMatch(home, away)
}

type League struct {
	Idable
	Name    string
	teamMap map[string]*Team
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
		teamMap:     teamMap,
		teams:       teams,
		Table:       NewTable(teams),
		Rounds:      rounds,
		totalRounds: (len(teams) * 2) - 2,
		RPointer:    0,
	}
}

func (l *League) IsFinished() bool {
	return l.RPointer >= l.totalRounds
}

func (l *League) NextRound() (*Round, bool) {
	if l.IsFinished() {
		return nil, false
	}
	return l.Rounds[l.RPointer].Round(l.teamMap), true
}

func (l *League) Update(round *Round) {

	l.Table.Update(round)
	//l. Update Stats
	l.RPointer++
}
