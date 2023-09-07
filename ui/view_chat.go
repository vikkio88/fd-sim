package ui

import (
	"fdsim/services"
	"fdsim/utils"
	"fdsim/vm"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func chatView(ctx *AppContext) *fyne.Container {
	game, _ := ctx.GetGameState()
	params := ctx.RouteParam.(vm.ChatParams)
	if params.IsSimpleChat() {
		return simpleChat(params, ctx)
	}

	hasTeam := params.Team != nil
	mI, _ := ctx.Db.MarketR().GetTransferMarketInfo()

	title := "Contract Offer"
	moneyLabel := "Yearly Wage:"
	if hasTeam {
		title = "Transfer Offer"
		moneyLabel = "Bid:"
	}

	value := params.ValueF
	bv := binding.BindFloat(&value)
	var contractYrsV binding.Float

	money := binding.NewString()
	moneySlider := widget.NewSliderWithData(value-(value*.5), 2*params.ValueF1, bv)
	moneySlider.Step = value * .1
	bv.AddListener(binding.NewDataListener(func() {
		v, _ := bv.Get()
		money.Set(utils.NewEurosFromF(v).StringKMB())
	}))

	playerInfo := container.NewVBox()
	if hasTeam {
		playerInfo.Add(
			container.NewGridWithColumns(2,
				valueLabel("Contract Yrs Left:", widget.NewLabel(fmt.Sprintf("%d", params.Player.YContract))),
				hL(params.Team.Name, func() { ctx.PushWithParam(TeamDetails, params.Team.Id) }),
			),
		)
	}
	playerInfo.Add(
		valueLabel("Role", widget.NewLabel(params.Player.Role.StringShort())))
	playerInfo.Add(
		valueLabel("Value", widget.NewLabel(vm.GetApproxMoney(params.Player.Value))))
	playerInfo.Add(
		valueLabel("Ideal Wage", widget.NewLabel(vm.GetApproxMoney(params.Player.IdealWage))))

	offerContent := container.NewVBox(
		makePlayerHeader(params.Player),
		playerInfo,
		h2("Your Offer"),
		valueLabel("Transfer Budget", widget.NewLabel(mI.TransferBudget.StringKMB())),
		valueLabel(moneyLabel, widget.NewLabelWithData(money)),
		moneySlider,
	)

	if params.ValueI != nil {
		cYearsStr := binding.NewString()
		y := float64(*params.ValueI)
		contractYrsV = binding.BindFloat(&y)
		yearsSlider := widget.NewSliderWithData(float64(1), float64(5), contractYrsV)
		yearsSlider.Step = 1

		contractYrsV.AddListener(binding.NewDataListener(
			func() {
				v, _ := contractYrsV.Get()
				cYearsStr.Set(fmt.Sprintf("%.0f", v))
			},
		))
		offerContent.Add(
			valueLabel("Contract Years:", widget.NewLabelWithData(cYearsStr)),
		)
		offerContent.Add(yearsSlider)
	}

	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(centered(h1(title))),
		).
		Bottom(rightAligned(widget.NewButton("Offer", func() {
			offer, _ := bv.Get()
			dialog.ShowConfirm("Making Offer", "Are you sure?", func(b bool) {
				if !b {
					ctx.Pop()
					return
				}

				// Offer
				if hasTeam {
					services.InstantDecisionBidForAPlayer(vm.MakeBidForPlayerParams(fdTeamId, params, offer), game, ctx.Db)
				} else {
					yf, _ := contractYrsV.Get()
					services.InstantDecisionOfferedContractToAPlayer(vm.MakePlayerContractOfferParams(fdTeamId, params, offer, int(yf)), game, ctx.Db)
				}

				ctx.Db.GameR().Update(game)
				ctx.BackToMain()

			}, ctx.GetWindow())
		}))).
		Get(
			offerContent,
		)

}

func simpleChat(params vm.ChatParams, ctx *AppContext) *fyne.Container {
	chatTitle := "Chat with: "
	if params.IsCoachChat() {
		chatTitle += fmt.Sprintf("%s (Coach)", params.Coach.String())
	} else {
		chatTitle += fmt.Sprintf("%s (Player)", params.Player.String())
	}
	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(centered(h1(chatTitle))),
		).
		Get(
			container.NewCenter(
				h1("Simple Chat"),
			),
		)
}
