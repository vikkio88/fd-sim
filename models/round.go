package models

import (
	"fdsim/libs"
	"fmt"

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
	WasPlayed bool
}

func NewRound(id string, index int, matches []*Match) *Round {
	mmap := map[string]*Match{}
	for _, m := range matches {
		mmap[m.Id] = m
	}
	return &Round{
		Idable:    NewIdable(id),
		Index:     index,
		Matches:   matches,
		MatchMap:  mmap,
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
