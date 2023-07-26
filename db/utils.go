package db

import (
	"fdsim/utils"
	"fmt"

	"github.com/oklog/ulid/v2"
)

func toMoney(value int64) utils.Money {
	return utils.NewEuros(value / moneyMultiplier)
}

func IdGenerator(suffix string) string {
	return fmt.Sprintf("%s_%s", suffix, ulid.Make())
}
