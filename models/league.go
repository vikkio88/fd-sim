package models

// Round Placeholder
type RPH struct {
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
			Index:   round,
			Matches: tempRoundFirstHalf,
		})

		secondHalf = append(secondHalf, RPH{
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
	return NewRound(r.Index, matches)
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
	Name    string
	teamMap map[string]*Team
	teams   []*Team
	// Rounds Placeholders
	Rounds []RPH
	Table  *Table
	// Last round pointer
	rPointer    int
	totalRounds int
}

func NewLeague(name string, teams []*Team) League {
	teamMap := map[string]*Team{}
	teamIds := make([]string, len(teams))
	for i, t := range teams {
		teamMap[t.Id] = t
		teamIds[i] = t.Id
	}
	rounds := NewRoundsCalendar(teamIds)
	return League{
		Name:        name,
		teamMap:     teamMap,
		teams:       teams,
		Table:       NewTable(teams),
		Rounds:      rounds,
		totalRounds: (len(teams) * 2) - 2,
		rPointer:    -1,
	}
}

func (l *League) NextRound() *Round {
	return nil
}
