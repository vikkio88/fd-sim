package db

import "fdsim/utils"

func toMoney(value int64) utils.Money {
	return utils.NewEuros(value / moneyMultiplier)
}
