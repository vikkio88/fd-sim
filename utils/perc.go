package utils

import "fmt"

type Perc struct {
	val int
}

func NewPerc(val int) Perc {
	if val < 0 {
		val = 0
	}

	if val > 100 {
		val = 100
	}

	return Perc{val}
}

func (p *Perc) Val() int {
	return p.val
}

func (p Perc) Less(o Perc) bool {
	return p.val < o.val
}

func (p Perc) String() string {
	return fmt.Sprintf("%d%%", p.val)
}
