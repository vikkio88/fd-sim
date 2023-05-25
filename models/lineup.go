package models

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
