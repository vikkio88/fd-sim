package models

import (
	"fdsim/libs"
)

type Lineup struct {
	Module       Module
	Starting     map[Role][]PPH
	teamStats    TeamStats
	lineupStats  TeamStats
	SectorStat   map[Role]TeamStats
	starterCount int
	FlatPlayers  []PPH

	// Bench    map[Role][]PPH
	//TODO: track substitutions
}

func NewLineup(module Module, starting map[Role][]PPH, stats TeamStats) *Lineup {
	starterCount := 0
	sectorStat := map[Role]TeamStats{}
	flattened := []PPH{}
	for r, ps := range starting {
		starterCount += len(ps)
		s, m, a := calculateAvgs(ps)
		sectorStat[r] = TeamStats{s, m, a}
		flattened = append(flattened, ps...)
	}
	s, m, a := calculateAvgs(flattened)
	lineupStats := TeamStats{s, m, a}

	return &Lineup{
		Module:       module,
		Starting:     starting,
		teamStats:    stats,
		lineupStats:  lineupStats,
		starterCount: starterCount,
		SectorStat:   sectorStat,
		FlatPlayers:  flattened,
		//TODO: calculate also starting skillsAvg

		// TODO: model issues like missing players in role or similar
		// Bench:    bench,
	}
}

func (l *Lineup) Ids() []string {
	ids := make([]string, len(l.FlatPlayers))
	for i, p := range l.FlatPlayers {
		ids[i] = p.Id
	}
	return ids
}

func (l *Lineup) CountStarters() int {
	return l.starterCount
}

func (l *Lineup) Malus() (int, int) {
	bonusOpponent := 0
	malusSelf := 0
	if !l.Module.Validate(l.Starting) {
		bonusOpponent += 1
	}

	if l.CountStarters() != 11 {
		bonusOpponent += 3
		malusSelf += 11 - l.starterCount
	}

	return bonusOpponent, malusSelf
}

func (l *Lineup) BestPlayerInRole(role Role) (*PPH, bool) {
	pls, ok := l.Starting[role]
	if !ok || len(pls) < 1 {
		return nil, false
	}

	return &pls[0], true
}

func (l *Lineup) Scorers(count int, rng *libs.Rng) []string {
	scorers := []string{}
	for i := 0; i < count; i++ {
		scorers = append(scorers, l.Scorer(rng))
	}
	return scorers
}

func (l *Lineup) Scorer(rng *libs.Rng) string {
	role := MF
	if len(l.Starting[MF]) < 1 {
		role = ST
	}

	if rng.ChanceI(70) {
		role = ST
	} else if rng.ChanceI(30) {
		role = DF
	}

	if role == ST && len(l.Starting[ST]) < 1 {
		role = DF
	}

	if len(l.Starting[DF]) < 1 {
		role = GK
	}

	if role == GK && len(l.Starting[GK]) < 1 {
		panic("There are no players in this lineup, this should never happen")
	}

	Idx := rng.Index(len(l.Starting[role]))
	return l.Starting[role][Idx].Id
}

func calculateAvgs(players []PPH) (float64, float64, float64) {
	tot := len(players)
	totS := 0
	totM := 0
	totA := 0
	for _, p := range players {
		totS += p.Skill
		totM += p.Morale
		totA += p.Age
	}

	valS := float64(totS) / float64(tot)
	valM := float64(totM) / float64(tot)
	valA := float64(totA) / float64(tot)
	return valS, valM, valA
}

func (l *Lineup) Score(owngoals, othergoals int, scorers []string, rng *libs.Rng) PlayerScoreMap {
	res := PlayerScoreMap{}
	min, max := 40, 100
	if owngoals > othergoals {
		min, max = 60, 100
	} else if owngoals < othergoals {
		min, max = 20, 60
	}

	for _, p := range l.FlatPlayers {
		pId := p.Id
		if isAScorer(pId, scorers) {
			min += 15
		}

		if p.Morale > 50 && rng.ChanceI(p.Morale) {
			min += 10
		}

		if p.Morale < 50 {
			min -= 15
		}

		if min <= 10 {
			// adjusting so the min score is not that low
			min = 20
		}

		score := float64(rng.UInt(min, max)) / 10.
		res[pId] = score
	}

	return res
}

func isAScorer(p string, scorers []string) bool {
	for _, s := range scorers {
		if p == s {
			return true
		}
	}

	return false
}
