package widgets

import (
	"fdsim/res"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Icon(name string) *widget.Icon {
	switch name {
	//TODO: move those to enums
	case "city":
		return widget.NewIcon(theme.NewThemedResource(res.City))
	case "dumbell":
		return widget.NewIcon(theme.NewThemedResource(res.Dumbell))
	case "team":
		return widget.NewIcon(theme.NewThemedResource(res.Team))
	case "money":
		return widget.NewIcon(theme.NewThemedResource(res.Money))
	case "sad_face":
		return widget.NewIcon(theme.NewThemedResource(res.SadFace))
	case "meh_face":
		return widget.NewIcon(theme.NewThemedResource(res.MehFace))
	case "happy_face":
		return widget.NewIcon(theme.NewThemedResource(res.HappyFace))
	case "transfers":
		return widget.NewIcon(theme.NewThemedResource(res.Transfers))
	case "email_read":
		return widget.NewIcon(theme.NewThemedResource(res.EmailRead))
	case "contract":
		//TODO: check how to theme also icons that colour changes if themed
		return widget.NewIcon(res.Contract)
	case "newspaper":
		return widget.NewIcon(res.NewsPaper)
	case "newspaper_read":
		return widget.NewIcon(res.NewsPaperRead)

	}

	return widget.NewIcon(theme.NewThemedResource(res.Unknown))
}
