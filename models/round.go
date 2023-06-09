package models

import (
	"fdsim/libs"
	"fdsim/utils"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

const roundInMemoryId = "rdId"

func roundIdGenerator() string {
	return fmt.Sprintf("%s_%s", roundInMemoryId, ulid.Make())
}

type RoundResult struct {
	Id        string
	Index     int
	Matches   []*MatchResult
	WasPlayed bool
}

type Round struct {
	Idable
	Index     int
	Matches   []*Match
	MatchMap  map[string]*Match
	resultMap map[string]*Result
	Date      time.Time
	WasPlayed bool
}

func NewRound(id string, index int, date time.Time, matches []*Match) *Round {
	mmap := map[string]*Match{}
	for _, m := range matches {
		mmap[m.Id] = m
	}
	return &Round{
		Idable:    NewIdable(id),
		Index:     index,
		Matches:   matches,
		MatchMap:  mmap,
		Date:      date,
		resultMap: map[string]*Result{},
		WasPlayed: false,
	}
}

func (r *Round) Simulate(rng *libs.Rng) {
	for _, m := range r.Matches {
		m.Simulate(rng)
		if res, ok := m.Result(); ok {
			r.resultMap[m.Id] = res
		}
	}

	r.WasPlayed = true
}

func (r *Round) Results() (map[string]*Result, bool) {
	if !r.WasPlayed {
		return nil, false
	}

	return r.resultMap, true
}

// Round PH With Team PH
type RPHTPH struct {
	Id      string
	Index   int
	Played  bool
	Date    time.Time
	Matches []MPHTPH
}

// Round Placeholder
type RPH struct {
	Id      string
	Index   int
	Date    time.Time
	Matches []MPH
}

func (r *RPH) Round(teamsMap TeamMap) *Round {
	matches := make([]*Match, len(r.Matches))
	for i, mph := range r.Matches {
		home, ok := teamsMap[mph.Home]
		away, ok2 := teamsMap[mph.Away]
		if !ok || !ok2 {
			panic("Empty team map")
		}
		matches[i] = mph.Match(mph.Id, home, away)

	}
	return NewRound(r.Id, r.Index, r.Date, matches)
}

func (r *RPH) RoundTPH(teamsMap TeamMap) *RPHTPH {
	matches := make([]MPHTPH, len(r.Matches))
	for i, mph := range r.Matches {
		home, ok := teamsMap[mph.Home]
		away, ok2 := teamsMap[mph.Away]
		if !ok || !ok2 {
			panic("Empty team map")
		}
		matches[i] = *mph.MPHTPH(mph.Id, home.PH(), away.PH())

	}
	return &RPHTPH{Id: r.Id, Index: r.Index, Matches: matches, Date: r.Date}
}

func NewRoundsCalendar(teamIds []string, seasonStartYear int) []RPH {
	var firstHalf []RPH
	var secondHalf []RPH
	numberOfTeams := len(teamIds)
	initialMonth := time.September
	if numberOfTeams > 18 {
		initialMonth = time.August
	}
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
			Date:    GetRoundDateByIndex(seasonStartYear, initialMonth, round, false),
			Matches: tempRoundFirstHalf,
		})

		secondHalf = append(secondHalf, RPH{
			Id:    roundIdGenerator(),
			Index: round + totalRounds,
			// here is round and not round+total because I start anew from Jan
			Date:    GetRoundDateByIndex(seasonStartYear, initialMonth, round, true),
			Matches: tempRoundSecondHalf,
		})
	}

	firstHalf = append(firstHalf, secondHalf...)

	return firstHalf
}

func GetRoundDateByIndex(initialYear int, initialMonth time.Month, index int, secondHalf bool) time.Time {
	offset := 7 * index
	if initialMonth == time.August && !secondHalf {
		return utils.GetLastSunday(initialYear, initialMonth).AddDate(0, 0, offset)
	}

	if !secondHalf {
		return utils.GetFirstSunday(initialYear, initialMonth).AddDate(0, 0, offset)
	}

	nextYear := initialYear + 1
	firstSundayOfJan := utils.GetFirstSunday(nextYear, time.January)
	return firstSundayOfJan.AddDate(0, 0, 7+offset)
}
