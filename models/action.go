package models

import "time"

type Action struct {
	Date        time.Time
	Description string
	Decision    Decision
}

type Decision struct {
	YN    *bool
	Value *float64
}

func NewAction(description string, decision Decision, date time.Time) Action {
	return Action{
		Date:        date,
		Description: description,
		Decision:    decision,
	}
}
