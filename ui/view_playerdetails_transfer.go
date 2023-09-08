package ui

import (
	"fdsim/conf"
	"fdsim/models"
	"fdsim/services"
	vm "fdsim/vm"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func makePTransferTab(ctx *AppContext, player *models.PlayerDetailed, canSeeDetails bool) fyne.CanvasObject {
	g, _ := ctx.GetGameState()
	tInfo, ok := ctx.Db.MarketR().GetTransferMarketInfo()
	var offer *models.Offer

	if !ok {
		// this should not happen as it wont appear if you have no team
		// the player view actually prevents this from being rendered so it should never happen
		panic("you should not see this if you are hired")
	}

	fdTeam := g.GetTeamOrEmpty()
	teamAcceptedOffer := false
	readyToTransfer := false
	waitingForResponse := false
	status := centered(h3(""))
	decisionKey := player.Id
	canCancel := false

	var actionBtn *widget.Button
	cancelBtn := widget.NewButton("Cancel", func() {})
	cancelBtn.Disable()

	if o, ok := player.GetOfferFromTeamId(fdTeam.Id); ok {
		decisionKey = fmt.Sprintf("%s.%s", decisionKey, o.OfferingTeam.Id)
		offer = o
		cancelBtn.OnTapped = func() {
			dialog.ShowConfirm("Withdraw Offer", "Are you sure?", func(b bool) {
				if !b {
					return
				}
				services.InstantDecisionCancelOffer(o, ctx.Db)
				ctx.BackToMain()
			}, ctx.GetWindow())
		}

		offerDateStr := o.LastUpdate.Format(conf.DateFormatShort)

		switch o.Stage() {
		case models.OfstOffered:
			waitingForResponse = true
			canCancel = true
			status = centered(h3(
				fmt.Sprintf("Your made an offer for this player on %s (%s).", offerDateStr, o.BidValue.StringKMB()),
			))
		case models.OfstContractOffered:
			waitingForResponse = true
			canCancel = true
			status = centered(h3(
				fmt.Sprintf("Your made a contract offer for this player on %s (%s / %d years).", offerDateStr, o.WageValue.StringKMB(), *o.YContract),
			))
		case models.OfstTeamAccepted:
			teamAcceptedOffer = true
			canCancel = true
			status = centered(h3(
				fmt.Sprintf("%s accepted your %s offer on %s.", o.Team.Name, o.BidValue.StringKMB(), offerDateStr),
			))
			//TODO: add info here
		case models.OfstReady:
			readyToTransfer = true
			canCancel = true
			status = centered(h3(
				fmt.Sprintf("This player accepted your contract offer on %s (%s / %d years).", offerDateStr, o.WageValue.StringKMB(), *o.YContract),
			))
		case models.OfstReadyTP:
			readyToTransfer = true
			canCancel = true
			status = centered(container.NewVBox(
				h3(fmt.Sprintf("%s accepted your %s offer.", o.Team.Name, o.BidValue.StringKMB())),
				h3(
					fmt.Sprintf(
						"This player accepted your contract offer of %s for %d years on the %s",
						o.WageValue.StringKMB(),
						*o.YContract,
						offerDateStr,
					),
				)))
		}
	}
	tV := vm.NewApproxTransferVals(player)

	if readyToTransfer {
		actionBtn = widget.NewButton("Confirm", func() {
			dialog.NewConfirm("Confirming Transfer", "Are you sure?", func(b bool) {
				if !b {
					return
				}
				services.InstantDecisionConfirmInTransfer(offer, g, ctx.Db)
				ctx.BackToMain()

			}, ctx.GetWindow())
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

	if canCancel {
		cancelBtn.Enable()
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
		Bottom(rightAligned(container.NewHBox(cancelBtn, actionBtn))).
		Get(info)
}
