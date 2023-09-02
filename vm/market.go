package vm

import (
	"fdsim/models"
	"fdsim/utils"
	"fmt"
)

// return iWage, wage, value, lowerV, higherV, lowerW, higherW, isFreeAgent
func GetApproxMoney(money utils.Money) string {
	valueLow, valueHigh := utils.GetApproxRangeM(money)
	return fmt.Sprintf("Value: %s - %s", valueLow.StringKMB(), valueHigh.StringKMB())
}

type ApproxTransferVals struct {
	IWage       string
	Wage        string
	Value       string
	LowerV      float64
	HigherV     float64
	LowerW      float64
	HigherW     float64
	IsFreeAgent bool
}

func NewApproxTransferVals(player *models.PlayerDetailed) ApproxTransferVals {
	iWage := GetApproxMoney(player.IdealWage)
	wage := GetApproxMoney(player.Wage)
	value := GetApproxMoney(player.Value)

	lowerV, higherV := utils.GetApproxRangeF(player.Value.Value())
	lowerW, higherW := utils.GetApproxRangeF(player.IdealWage.Value())

	isFreeAgent := player.Team == nil
	return ApproxTransferVals{iWage, wage, value, lowerV, higherV, lowerW, higherW, isFreeAgent}
}
