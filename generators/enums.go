package generators

import (
	"fdsim/enums"
	"fdsim/libs"
)

type EnumsGen struct {
	rng       *libs.Rng
	countries []enums.Country
}

func NewEnumsGen(seed int64) *EnumsGen {
	rng := libs.NewRng(seed)
	return NewEnumsGenSeeded(rng)
}

func NewEnumsGenSeeded(rng *libs.Rng) *EnumsGen {
	return &EnumsGen{rng, enums.AllCountries()}
}

func (e *EnumsGen) Country() enums.Country {
	idx := e.rng.Index(len(e.countries))

	return e.countries[idx]
}
