package utils

import "fmt"

type Perc struct {
	val int
}

func cap(val, min, max int) int {
	if val < min {
		val = min
	}

	if val > max {
		val = max
	}

	return val
}

func NewPerc(val int) Perc {
	val = cap(val, 0, 100)
	return Perc{val}
}

func (p *Perc) SetVal(val int) {
	val = cap(val, 0, 100)
	p.val = val

}
func (p *Perc) Val() int {
	return p.val
}

func (p Perc) Less(o Perc) bool {
	return p.LessThan(o.val)
}

func (p Perc) LessThan(o int) bool {
	return p.val < o
}

func (p Perc) String() string {
	return fmt.Sprintf("%d%%", p.val)
}
