package models

import "fmt"

type Award struct {
	//TODO Maybe to a LPH?
	LeagueId   string
	LeagueName string

	Scorer bool
	Mvp    bool
	Score  float64
	Goals  int
	Played int
	Team   TPH
}

func (a *Award) StatString() string {

	if a.Mvp && a.Scorer {
		score := "-"
		if a.Played > 0 {
			score = fmt.Sprintf("%.2f", a.Score/float64(a.Played))
		}

		return fmt.Sprintf("%d (%d) s: %s", a.Goals, a.Played, score)
	}

	if a.Mvp {
		score := "-"
		if a.Played > 0 {
			score = fmt.Sprintf("%.2f", a.Score/float64(a.Played))
		}
		return fmt.Sprintf("%s (%d)", score, a.Played)
	}

	return fmt.Sprintf("%d (%d)", a.Goals, a.Played)
}

func (a *Award) String() string {
	awards := ""
	if a.Mvp {
		awards += "MVP"
	}

	if a.Scorer {
		if awards != "" {
			awards += ", Top Scorer"
		} else {
			awards += "Top Scorer"
		}
	}
	return awards
}

type Trophy struct {
	//TODO Maybe to a LPH?
	LeagueId   string
	LeagueName string

	Team TPH
	Year int
}
