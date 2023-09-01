package generators

import (
	"fdsim/data"
	"fdsim/libs"
)

type RngDataSlice []string

func (s RngDataSlice) One() string {
	return s[libs.NewRngAutoSeeded().Index(len(s))]
}

func (s RngDataSlice) OneSeed(seed int64) string {
	rng := libs.NewRng(seed)
	return s[rng.Index(len(s))]
}

var EmailDomains RngDataSlice = data.EmailDomains
