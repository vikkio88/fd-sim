package models

import "time"

type ActionType uint8

const (
	ActionRespondContract ActionType = iota
	ActionOutTranfer
	ActionInTranfer
	ActionPlayerContract

	ActionTest

	Blank
)

type Actionable struct {
	Description string
	ActionType  ActionType
	Expires     time.Time
	Choices     Choosable
	Decision    *Choosable
}

type Choosable struct {
	ActionType ActionType
	YN         *bool
	ValueInt   *int
	ValueF     *float64
	Label      *string
	PlayerId   *string
	TeamId     *string
	Item1      *string
	Item2      *string
	Item3      *string
	Item4      *string
}

func NewActionable(description string, choices Choosable, date time.Time, actionType ActionType) *Actionable {
	return &Actionable{
		ActionType:  actionType,
		Expires:     date,
		Description: description,
		Choices:     choices,
	}
}

func (a *Actionable) setDecision() {
	if a.Decision == nil {
		a.Decision = &Choosable{}
	}
}

func (a *Actionable) Decide(decision *Choosable) {
	a.Decision = decision
}

func (a *Actionable) AnswerYN(yn *bool) {
	a.setDecision()
	a.Decision.YN = yn
}

func (a *Actionable) AnswerValue(value *float64) {
	a.setDecision()
	a.Decision.ValueF = value
}
