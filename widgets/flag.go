package widgets

import (
	"fdsim/enums"
	"fdsim/res"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FlagIcon(country enums.Country) *widget.Icon {
	switch country {
	case enums.IT:
		return widget.NewIcon(res.It)
	case enums.EN:
		return widget.NewIcon(res.En)
	case enums.FR:
		return widget.NewIcon(res.Fr)
	case enums.ES:
		return widget.NewIcon(res.Es)
	case enums.DE:
		return widget.NewIcon(res.De)
	}
	return widget.NewIcon(res.It)
}

type Flag struct {
	widget.BaseWidget
	country   enums.Country
	container *fyne.Container
}

func NewFlag(country enums.Country) *Flag {
	f := &Flag{country: country}
	f.ExtendBaseWidget(f)
	return f
}

func (f *Flag) Refresh() {
	f.update()
	f.BaseWidget.Refresh()
}

func (f *Flag) SetCountry(country enums.Country) {
	f.country = country
	f.Refresh()
}

func (f *Flag) update() {
	if f.container == nil {
		f.container = container.NewCenter()
	} else {
		f.container.RemoveAll()
	}

	f.container.Add(FlagIcon(f.country))
}

func (f *Flag) CreateRenderer() fyne.WidgetRenderer {
	f.update()
	return widget.NewSimpleRenderer(
		f.container,
	)
}
