package libs

import (
	"fdsim/utils"
	"math/rand"
)

const maxPerc = 100

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

func (r *Rng) Index(len int) int {
	return r.UInt(0, len-1)
}

func (r *Rng) UInt(min, max int) int {
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

// Need generics for this and fyne does not support that version of go yet
// func (r *Rng) PickOne(list []any) any {
// 	i := r.UInt(0, len(list))

// 	return list[i]
// }
