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

	if offer, ok := player.GetOfferFromTeamId(fdTeam.Id); ok {
		decisionKey = fmt.Sprintf("%s.%s", decisionKey, offer.OfferingTeam.Id)
		cancelBtn.OnTapped = func() {
			dialog.ShowConfirm("Withdraw Offer", "Are you sure?", func(b bool) {
				if !b {
					return
				}
				services.InstantDecisionCancelOffer(offer, ctx.Db)
				ctx.BackToMain()
			}, ctx.GetWindow())
		}

		offerDateStr := offer.LastUpdate.Format(conf.DateFormatShort)

		switch offer.Stage() {
		case models.OfstOffered:
			waitingForResponse = true
			canCancel = true
			status = centered(h3(
				fmt.Sprintf("Your made an offer for this player on %s (%s).", offerDateStr, offer.BidValue.StringKMB()),
			))
		case models.OfstContractOffered:
			waitingForResponse = true
			canCancel = true
			status = centered(h3(
				fmt.Sprintf("Your made a contract offer for this player on %s (%s / %d years).", offerDateStr, offer.WageValue.StringKMB(), *offer.YContract),
			))
		case models.OfstTeamAccepted:
			teamAcceptedOffer = true
			canCancel = true
			status = centered(h3(
				fmt.Sprintf("%s accepted your %s offer on %s.", offer.Team.Name, offer.BidValue.StringKMB(), offerDateStr),
			))
			//TODO: add info here
		case models.OfstReady:
			readyToTransfer = true
			canCancel = true
			status = centered(h3(
				fmt.Sprintf("This player accepted your contract offer on %s (%s / %d years).", offerDateStr, offer.WageValue.StringKMB(), *offer.YContract),
			))
		case models.OfstReadyTP:
			readyToTransfer = true
			canCancel = true
			status = centered(container.NewVBox(
				h3(fmt.Sprintf("%s accepted your %s offer.", offer.Team.Name, offer.BidValue.StringKMB())),
				h3(
					fmt.Sprintf(
						"This player accepted your contract offer of %s for %d years on the %s",
						offer.WageValue.StringKMB(),
						*offer.YContract,
						offerDateStr,
					),
				)))
		}
	}
	tV := vm.NewApproxTransferVals(player)

	if readyToTransfer {
		actionBtn = widget.NewButton("Confirm", func() {
			ep := models.EP()
			ep.PlayerId = player.Id
			ep.PlayerName = player.String()
			ep.FdTeamId = fdTeam.Id
			ep.FdTeamName = fdTeam.Name
			decision := vm.MakeAcceptTransferDecision(g, ep)
			g.QueueDecision(decision)
			addPendingDecision(decision.Choice.ActionType, decisionKey)
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

	if pendingDecision.Has(models.ActionConfirmInTranfer, decisionKey) {
		actionBtn.Disable()
		actionBtn.SetText("Accepted")
		cancelBtn.Disable()
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
