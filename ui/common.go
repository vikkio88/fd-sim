package ui

import (
	"fdsim/conf"
	"fdsim/utils"
	vm "fdsim/vm"
	"fdsim/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func h1(text string) *canvas.Text {
	txt := canvas.NewText(text, theme.ForegroundColor())
	txt.TextSize = 20
	txt.Alignment = fyne.TextAlignCenter
	return txt
}

// h1 aligned Right
func h1R(text string) *canvas.Text {
	txt := h1(text)
	txt.Alignment = fyne.TextAlignTrailing
	return txt
}

// h1 aligned Left
func h1L(text string) *canvas.Text {
	txt := h1(text)
	txt.Alignment = fyne.TextAlignLeading
	return txt
}

func h2(text string) *canvas.Text {
	txt := canvas.NewText(text, theme.ForegroundColor())
	txt.TextSize = 18
	return txt
}

func small(text string) *canvas.Text {
	txt := canvas.NewText(text, theme.ForegroundColor())
	txt.TextSize = 10
	return txt
}

func centered(object fyne.CanvasObject) *fyne.Container {
	return container.NewCenter(object)
}
func rightAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, nil, object)
}

func leftAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, object, nil)
}

func topAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(object, nil, nil, nil)
}

func bottomAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, object, nil, nil)
}

func topNavBar(ctx *AppContext) fyne.CanvasObject {
	game, ok := ctx.GetGameState()
	nav := container.NewHBox(
		widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			ctx.Pop()
		}),
		widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
			ctx.BackToMain()
		}),
	)

	if ok {
		nav.Add(widget.NewLabel(fmt.Sprintf("%s", game.Date.Format(conf.DateFormatGame))))
	}

	return nav
}

func starsFromPerc(perc utils.Perc) fyne.CanvasObject {
	return widgets.NewStarRatingFromFloat(vm.PercToStars(perc))
}

func starsFromf64(value float64) fyne.CanvasObject {
	return widgets.NewStarRatingFromFloat(vm.PercFToStars(value))
}

func valueLabel(label string, value fyne.CanvasObject) *fyne.Container {
	labelLbl := boldLabel(label)
	return container.NewGridWithColumns(2,
		centered(labelLbl),
		value,
	)
}
func boldLabel(label string) *widget.Label {
	labelLbl := widget.NewLabel(label)
	labelLbl.TextStyle = fyne.TextStyle{Bold: true}

	return labelLbl
}

func signalFdTeamTxt(txt string) string {
	return fmt.Sprintf("[ %s ]", txt)
}

func signalFdTeam(lbl *widget.Label) {
	lbl.SetText(signalFdTeamTxt(lbl.Text))
	lbl.TextStyle = fyne.TextStyle{Bold: true}
}

// Little utility to get a Centered HL from row in a list
// had to force center() as otherwise the underline goes all over the space
func getCenteredHL(o fyne.CanvasObject) *widget.Hyperlink {
	return o.(*fyne.Container).Objects[0].(*widget.Hyperlink)
}

// Returns a Centered HL
func hL(label string, onTapped func()) fyne.CanvasObject {
	hl := widget.NewHyperlink(label, nil)
	hl.OnTapped = onTapped
	return centered(hl)
}

// return an approximation of money
func getApproxMoney(money utils.Money) string {
	valueLow, valueHigh := utils.GetApproxRangeM(money)
	return fmt.Sprintf("Value: %s - %s", valueLow.StringKMB(), valueHigh.StringKMB())
}

// handle subtab navigation
func handleSubtabs(subtabIndex int, tabContainer *container.AppTabs) {
	if subtabIndex != -1 {
		if subtabIndex > len(tabContainer.Items) {
			return
		}
		tabContainer.SelectIndex(subtabIndex)
	}
}
