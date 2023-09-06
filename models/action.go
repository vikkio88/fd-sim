package models

import "time"

type ActionType uint8

const (
	ActionRespondContract ActionType = iota
	ActionConfirmOutTranfer
	ActionPlayerContract

	// FD Offered Contract to Player
	ActionPlayerContractOffer
	// FD Offered Money for a Player to a Team
	ActionPlayerOffer
	// FD confirmed the Transfer
	ActionConfirmInTranfer
	// FD Cancels Transfer
	ActionCancelTransfer

	ActionTest

	Blank
)

func (at ActionType) Choosable(params EventParams) Choosable {
	switch at {
	case ActionRespondContract:
		{
			var yn bool
			return Choosable{
				ActionType: at,
				YN:         &yn,
				Params:     params,
			}
		}
	}

	return Choosable{
		ActionType: at,
		Params:     params,
	}
}

func (at ActionType) Actionable(date time.Time, params EventParams) *Actionable {
	switch at {
	case ActionRespondContract:
		{

			return NewActionable(
				at,
				"Contract Offer",
				at.Choosable(params),
			)
		}

		//TODO: Remove Testing Action
	case ActionTest:
		var yn bool
		return NewActionable(
			at,
			"Testing Actionables",
			Choosable{
				ActionType: at,
				YN:         &yn,
				Params:     params,
			},
		)
	}

	return nil
}

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

func NewActionable(actionType ActionType, description string, choices Choosable) *Actionable {
	return &Actionable{
		ActionType:  actionType,
		Description: description,
		Choices:     choices,
	}
}
