package viewmodels

import "fyne.io/fyne/v2/data/binding"

func ClearDataUtList(dataList binding.UntypedList) {
	l, _ := dataList.Get()
	l = l[:0]
	dataList.Set(l)
}
