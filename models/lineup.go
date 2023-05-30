package models

import "fdsim/libs"

type Lineup struct {
	Module    Module
	Starting  map[Role][]PPH
	teamStats TeamStats
	// Bench    map[Role][]PPH
	//TODO: track substitutions
}

func NewLineup(module Module, starting map[Role][]PPH, stats TeamStats) *Lineup {
	return &Lineup{
		Module:    module,
		Starting:  starting,
		teamStats: stats,
		//TODO: calculate also starting skillsAvg

		// TODO: model issues like missing players in role or similar
		// Bench:    bench,
	}
}

func (l *Lineup) CountStarters() int {
	c := 0
	for _, ps := range l.Starting {
		c += len(ps)
	}

	return c
}

func (l *Lineup) Malus() (int, int) {
	malusOpponent := 0
	malusSelf := 0
	if !l.Module.Validate(l.Starting) {
		malusOpponent += 1
	}

	if l.CountStarters() != 11 {
		malusOpponent += 3
		malusSelf += 3
	}

	return malusOpponent, malusSelf
}

func (l *Lineup) BestPlayerInRole(role Role) (*PPH, bool) {
	pls, ok := l.Starting[role]
	if !ok || len(pls) < 1 {
		return nil, false
	}

	return &pls[0], true
}

func (l *Lineup) Scorer(rng *libs.Rng) string {
	role := MF
	if rng.ChanceI(70) {
		role = ST
	} else if rng.ChanceI(30) {
		role = DF
	}

	Idx := rng.Index(len(l.Starting[role]))
	return l.Starting[role][Idx].Id
}
