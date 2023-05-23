package utils

import (
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
