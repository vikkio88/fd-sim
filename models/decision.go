package models

import "time"

type Decision struct {
	Choice  Choosable
	EmailId *string
	Date    time.Time
}

func NewDecisionFromEmail(date time.Time, choice Choosable, emailId string) *Decision {
	return &Decision{
		Choice:  choice,
		EmailId: &emailId,
		Date:    date,
	}
}

func NewDecision(date time.Time, choice Choosable) *Decision {
	return &Decision{
		Choice: choice,
		Date:   date,
	}
}
