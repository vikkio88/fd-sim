package models

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
	Choices     Choosable
}

// This has to be in sync with services/parameters.go
type Choosable struct {
	ActionType ActionType
	YN         *bool
	Params     EventParams
}

func NewActionable(description string, choices Choosable, actionType ActionType) *Actionable {
	return &Actionable{
		ActionType:  actionType,
		Description: description,
		Choices:     choices,
	}
}
