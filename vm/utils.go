package vm

import (
	"fdsim/utils"

	"fyne.io/fyne/v2/data/binding"
)

func ClearDataUtList(dataList binding.UntypedList) {
	l, _ := dataList.Get()
	l = l[:0]
	dataList.Set(l)
}

func PercToStars(perc utils.Perc) float32 {
	val := float32(perc.Val()) * 5. / 100.
	return val
}

func PercFToStars(value float64) float32 {
	val := float32(value) * 5. / 100.
	return val
}
