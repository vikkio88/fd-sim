package models

import (
	"fdsim/utils"
	"time"
)

type TransferFailReason uint8

const (
	// No Reason To fail
	TFRNone TransferFailReason = iota
	TFRInsufficientFunds
)

type TransferResult struct {
	Success      bool
	Player       PNPH
	PreviousTeam *TPH
	Team         TPH
	Bid          utils.Money
	Wage         utils.Money
	YContract    int
	FailReason   TransferFailReason
	Date         time.Time
}

func NewTransferSuccess(player PNPH, team TPH, wage utils.Money, yContract int, date time.Time) *TransferResult {
	return &TransferResult{
		Success:    true,
		Player:     player,
		Team:       team,
		Wage:       wage,
		YContract:  yContract,
		FailReason: TFRNone,
		Date:       date,
	}
}

func NewTransferFail(player PNPH, team TPH) *TransferResult {
	return &TransferResult{
		Success:    false,
		Player:     player,
		Team:       team,
		FailReason: TFRInsufficientFunds,
	}
}

func (tr *TransferResult) AddTeam(previousTeam TPH, bid utils.Money) {
	tr.PreviousTeam = &previousTeam
	tr.Bid = bid
}
