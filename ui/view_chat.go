package ui

import (
	"fdsim/models"
	"fdsim/utils"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type ChatParams struct {
	IsPlayerOffer bool
	IsChat        bool
	Player        *models.PlayerDetailed
	Team          *models.TPH
	Coach         *models.Coach
	ValueF        float64
	ValueF1       float64
	ValueI        *int
}

func chatView(ctx *AppContext) *fyne.Container {
	params := ctx.RouteParam.(ChatParams)
	if params.IsChat && !params.IsPlayerOffer {
		return simpleChat(params, ctx)
	}

	hasTeam := params.Team != nil
	mI, _ := ctx.Db.GameR().GetTransferMarketInfo()

	title := "Contract Offer"
	moneyLabel := "Yearly Wage:"
	if hasTeam {
		title = "Transfer Offer"
		moneyLabel = "Bid:"
	}

	value := params.ValueF
	bv := binding.BindFloat(&value)
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
		valueLabel("Value", widget.NewLabel(getApproxMoney(params.Player.Value))))
	playerInfo.Add(
		valueLabel("Ideal Wage", widget.NewLabel(getApproxMoney(params.Player.IdealWage))))

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
		cyV := binding.BindFloat(&y)
		yearsSlider := widget.NewSliderWithData(float64(1), float64(5), cyV)
		yearsSlider.Step = 1

		cyV.AddListener(binding.NewDataListener(
			func() {
				v, _ := cyV.Get()
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
		Get(
			offerContent,
		)

}

func simpleChat(params ChatParams, ctx *AppContext) *fyne.Container {
	return NewFborder().
		Top(
			NewFborder().
				Left(topNavBar(ctx)).
				Get(centered(h1("Chat"))),
		).
		Get(
			container.NewCenter(
				h1("Simple Chat"),
			),
		)
}
