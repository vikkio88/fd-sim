package widgets

import (
	"fdsim/enums"
	"fdsim/res"

	"fyne.io/fyne/v2/widget"
)

func Flag(country enums.Country) *widget.Icon {
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
