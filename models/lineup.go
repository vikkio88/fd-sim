package models

import "fdsim/libs"

type Lineup struct {
	Module   Module
	Starting map[Role][]PPH
	// Bench    map[Role][]PPH
	//TODO: track substitutions
}

func NewLineup(module Module, starting map[Role][]PPH) Lineup {
	return Lineup{
		Module:   module,
		Starting: starting,

		// TODO: model issues like missing players in role or similar
		// Bench:    bench,
	}
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
