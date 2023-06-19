package models

import "fmt"

type Res12X uint8

const (
	R1 Res12X = iota
	R2
	RX
)

type ResultsPHMap map[string]*ResultPH

type ResultPH struct {
	MatchId   string
	GoalsHome int
	GoalsAway int
}

func (r *ResultPH) String() string {
	return fmt.Sprintf("%d - %d", r.GoalsHome, r.GoalsAway)

}

type Result struct {
	GoalsHome   int
	GoalsAway   int
	ScorersHome []string
	ScorersAway []string
}

func (r *Result) String() string {
	return fmt.Sprintf("%d - %d", r.GoalsHome, r.GoalsAway)

}

func NewResult(goalsHome, goalsAway int, scorersHome, scorersAway []string) *Result {
	return &Result{
		GoalsHome:   goalsHome,
		GoalsAway:   goalsAway,
		ScorersHome: scorersHome,
		ScorersAway: scorersAway,
	}
}

func (r *Result) X12() Res12X {
	switch {
	case r.Draw():
		return RX
	case r.HomeWin():
		return R1
	case r.AwayWin():
		return R2
	}

	//this should never happen
	panic("This should never happen")
}

func (r *Result) Draw() bool {
	return r.GoalsAway == r.GoalsHome
}

func (r *Result) HomeWin() bool {
	return r.GoalsHome > r.GoalsAway
}

func (r *Result) AwayWin() bool {
	return r.GoalsHome < r.GoalsAway
}
