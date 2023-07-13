package libs

import (
	"fdsim/utils"
	"math/rand"
)

const (
	maxPerc    = 100
	maxPercF   = 100.0
	percStdDev = 0.65
	percMean   = 0.13
)

type Rng struct {
	seed int64
	rand *rand.Rand
}

func NewRng(seed int64) *Rng {
	r := rand.New(rand.NewSource(seed))
	return &Rng{
		seed,
		r,
	}
}

func (r *Rng) PercR(min, max int) float64 {
	return float64(r.UInt(min, max)) / 100.
}

func (r *Rng) Perc() float64 {
	return float64(r.UInt(0, 100)) / 100.
}

func (r *Rng) NormF64(mean, stdDev float64) float64 {
	return rand.NormFloat64()*stdDev + mean
}

func (r *Rng) NormPercVal() int {
	return int(r.NormF64(percStdDev, percMean) * maxPercF)
}

func (r *Rng) Index(len int) int {
	return r.UInt(0, len-1)
}

func (r *Rng) UInt(min, max int) int {
	if min == max {
		return min
	}

	if min > max {
		max, min = min, max
	}

	if (max+1)-min < 0 || min < 0 {
		return 0
	}

	return r.rand.Intn((max+1)-min) + min
}

func (r *Rng) ChanceI(perc int) bool {
	p := utils.NewPerc(perc)
	return r.Chance(p)
}

func (r *Rng) Chance(perc utils.Perc) bool {
	return r.UInt(0, maxPerc) <= perc.Val()
}

func (r *Rng) PlusMinus(chance int) int {
	if r.ChanceI(chance) {
		return 1
	}
	return -1
}

func (r *Rng) PlusMinusVal(val, chance int) int {

	return r.PlusMinus(chance) * val
}

// Need generics for this and fyne does not support that version of go yet
// func (r *Rng) PickOne(list []any) any {
// 	i := r.UInt(0, len(list))

// 	return list[i]
// }
