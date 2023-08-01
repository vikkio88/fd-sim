package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const IntRangeSeparator = ".."

type IntRange struct {
	Min int
	Max int
}

func NewIntRange(min, max int) IntRange {
	if min > max {
		min, max = max, min
	}
	return IntRange{min, max}
}

func NewIntRangeFromStr(rng string) IntRange {
	pieces := strings.Split(rng, IntRangeSeparator)
	if len(pieces) < 2 {
		return NewIntRange(0, 0)
	}

	min, errm := strconv.ParseInt(pieces[0], 10, 0)
	max, errM := strconv.ParseInt(pieces[1], 10, 0)
	if errM != nil || errm != nil {
		return NewIntRange(0, 0)
	}

	return NewIntRange(int(min), int(max))
}

func GetApproxRangeI(number int) IntRange {
	power := int(math.Pow(10, float64(len(fmt.Sprint(number)))-1))
	lower := (number / power) * power
	upper := lower + power
	return NewIntRange(lower, upper)
}

func GetApproxRangeF(number float64) (float64, float64) {
	intPart, _ := math.Modf(number)
	power := math.Pow(10, math.Floor(math.Log10(math.Abs(intPart))))
	lower := math.Floor(intPart/power) * power
	upper := lower + power
	return lower, upper
}

func GetApproxRangeM(money Money) (Money, Money) {
	lower, upper := GetApproxRangeF(money.Value())

	return NewMoneyFromF(money.Currency, lower), NewMoneyFromF(money.Currency, upper)
}
