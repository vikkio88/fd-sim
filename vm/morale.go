package vm

import (
	"fdsim/utils"
)

type MoraleEmoj uint8

const (
	Happy MoraleEmoj = iota
	Sad
	Meh
)

const (
	happy string = "HAPPY"
	sad   string = "SAD"
	meh   string = "MEH"

	invalid string = "INVALID"
)

func getMapping() map[MoraleEmoj]string {
	return map[MoraleEmoj]string{
		Happy: happy,
		Meh:   meh,
		Sad:   sad,
	}
}

func getReverseMapping() map[string]MoraleEmoj {
	return map[string]MoraleEmoj{
		happy: Happy,
		sad:   Sad,
	}
}

func MoraleEmojFromPerc(perc utils.Perc) MoraleEmoj {
	if perc.Val() >= 80 {
		return Happy
	}

	if perc.Val() >= 50 {
		return Meh
	}

	return Sad
}

func (a MoraleEmoj) String() string {
	mapping := getMapping()
	if val, ok := mapping[a]; ok {
		return val
	}

	return invalid
}
