package ui

import (
	"fdsim/conf"
	"fdsim/models"
	vm "fdsim/vm"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makePTransferTab(ctx *AppContext, player *models.PlayerDetailed, canSeeDetails bool) fyne.CanvasObject {
	g, _ := ctx.GetGameState()
	teamAcceptedOffer := false
	readyToTransfer := false
	waitingForResponse := false
	status := centered(h2(""))

	if offer, ok := player.GetOfferFromTeamId(g.Team.Id); ok {
		switch offer.Stage() {
		case models.OfstOffered:
			waitingForResponse = true
			status = centered(h2(
				fmt.Sprintf("Your already made an offer for this player on %s (%s).", offer.OfferDate.Format(conf.DateFormatShort), offer.BidValue.StringKMB()),
			))
		case models.OfstContractOffered:
			waitingForResponse = true
			status = centered(h2(
				fmt.Sprintf("Your made a contract offer for this player (%s / %d years).", offer.WageValue.StringKMB(), *offer.YContract),
			))
		case models.OfstTeamAccepted:
			teamAcceptedOffer = true
			status = centered(h2(
				fmt.Sprintf("%s accepted your %s offer.", offer.Team.Name, offer.BidValue.StringKMB()),
			))
			//TODO: add info here
		case models.OfstReady:
		case models.OfstReadyTP:
			readyToTransfer = true
			status = centered(h2(
				fmt.Sprintf("Player Accepted the contract."),
			))
		}
	}

	tInfo, ok := ctx.Db.MarketR().GetTransferMarketInfo()

	if !ok {
		// this should not happen as it wont appear if you have no team
		panic("you should not see this if you are hired")
	}

	tV := vm.NewApproxTransferVals(player)
	var actionBtn *widget.Button
	if readyToTransfer {
		actionBtn = widget.NewButton("Confirm", func() {
			// make decision
		})
	} else if waitingForResponse {
		actionBtn = widget.NewButton("Waiting...", func() {})
		actionBtn.Disable()
	} else if !tV.IsFreeAgent && !teamAcceptedOffer {
		actionBtn = widget.NewButton("Make an Offer", func() {
			ctx.PushWithParam(Chat, vm.ChatParams{
				IsPlayerOffer: true,
				Player:        player,
				Team:          player.Team,
				ValueF:        tV.LowerV,
				ValueF1:       tV.HigherV,
			})
		})
	} else {
		actionBtn = widget.NewButton("Offer Contract", func() {
			contractY := 1
			ctx.PushWithParam(Chat, vm.ChatParams{
				IsPlayerOffer: true,
				Player:        player,
				ValueF:        tV.LowerW,
				ValueF1:       tV.HigherW,
				ValueI:        &contractY,
			})
		})
	}

	contractInfo := valueLabel("Contract", widget.NewLabel("-"))
	if !tV.IsFreeAgent {
		contractInfo = valueLabel("Contract", widget.NewLabel(fmt.Sprintf("%s / %d yrs", tV.Wage, player.YContract)))
	}

	info := container.NewVBox(
		status,
		valueLabel("Value: ", widget.NewLabel(tV.Value)),
		valueLabel("Ideal Wage: ", widget.NewLabel(tV.IWage)),
		contractInfo,
	)

	return NewFborder().Top(
		rightAligned(widget.NewLabel(fmt.Sprintf("Transfer Budget: %s", tInfo.TransferBudget.StringKMB()))),
	).
		Bottom(rightAligned(actionBtn)).
		Get(info)
}
