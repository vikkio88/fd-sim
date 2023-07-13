package ui

import (
	"fdsim/conf"
	"fdsim/db"
	"fdsim/models"
	"fdsim/vm"
	"fdsim/widgets"
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func notificationView(ctx *AppContext, route AppRoute) *fyne.Container {
	id := ctx.RouteParam.(string)
	var content fyne.CanvasObject
	var title string
	if route == Email {
		title = "Email"
		content = makeEmail(id, ctx)
	} else {
		title = "News"
		content = makeNews(id, ctx)
	}

	return NewFborder().
		Top(NewFborder().Left(topNavBar(ctx)).Get(centered(h1(title)))).
		Get(
			content,
		)
}

func makeNews(id string, ctx *AppContext) fyne.CanvasObject {
	news := ctx.Db.GameR().GetNewsById(id)
	return container.NewMax(
		widget.NewCard(
			news.Title,
			fmt.Sprintf("%s - %s", news.Date.Format(conf.DateFormatShort), news.NewsPaper),
			parseBody(news.Body, news.Links, ctx),
		),
	)
}

func makeEmail(id string, ctx *AppContext) fyne.CanvasObject {
	game, _ := ctx.GetGameState()
	email := ctx.Db.GameR().GetEmailById(id)

	body := NewFborder()
	content := container.NewVBox()
	emailBody := parseBody(email.Body, email.Links, ctx)
	content.Add(emailBody)

	if email.Action != nil {
		makeAction(email, content, body, game, ctx.Db)

	}
	return container.NewMax(
		widget.NewCard(
			email.Subject,
			fmt.Sprintf("%s - %s", email.Date.Format(conf.DateFormatShort), email.Sender),
			body.Get(content),
		),
	)
}

func makeAction(email *models.Email, content *fyne.Container, body *Fborder, game *models.Game, db db.IDb) {
	var actionable fyne.CanvasObject
	var replyBtn *widget.Button
	answered := email.Decision != nil
	answeredB := binding.NewBool()
	replyBtn = widget.NewButton("Reply", func() {
		answeredB.Set(true)
		//TODO: Set Decision and store email
		decision := email.Action.Choices
		email.Answer(&decision)

		decisionE := *email.Decision
		dec := models.NewDecisionFromEmail(game.Date, decisionE, email.Id)

		game.QueueDecision(dec)
		db.GameR().UpdateEmail(email)
		db.GameR().Update(game)

		// this forces the emails bind list to reload
		loadEmails(db)
	})

	if !answered {
		actionable = parseAction(email.Expires, email.Action)
		content.Add(actionable)
		body.Bottom(rightAligned(replyBtn))
		answeredB.AddListener(binding.NewDataListener(func() {
			if email.Decision == nil {
				return
			}

			actionable.Hide()
			content.Add(h1("You replied"))
			replyBtn.Disable()
		}))
	} else {
		content.Add(h1("You replied"))
	}
}

func parseAction(expires *time.Time, action *models.Actionable) fyne.CanvasObject {
	validity := ""
	if expires != nil && !expires.IsZero() {
		validity = fmt.Sprintf("valid until: %s", expires.Format(conf.DateFormatGame))
	}
	return container.NewVBox(
		widget.NewCard(
			action.Description,
			validity,
			makeChoices(action.Choices),
		),
	)
}

func makeChoices(choices models.Choosable) fyne.CanvasObject {
	c := container.NewVBox()

	if choices.YN != nil {
		yn := binding.BindBool(choices.YN)

		c.Add(
			// widget.NewCheckWithData("My Answer is YES", yn),
			widget.NewRadioGroup([]string{"Yes", "No"}, func(s string) {
				yn.Set(vm.YesNoToBool(s))
			}),
		)

		return c
	}

	// In case I want Value
	// if choices.Value != nil {
	// 	val := binding.BindFloat(choices.Value)
	// 	valueStr := ""
	// 	lbl := binding.BindString(&valueStr)
	// 	val.AddListener(binding.NewDataListener(
	// 		func() {
	// 			v, _ := val.Get()
	// 			lbl.Set(fmt.Sprintf("%s", utils.NewEurosFromF(v).StringKMB()))
	// 		},
	// 	))
	// 	slider := widget.NewSliderWithData(100, 200000000, val)
	// 	slider.Step = 1000
	// 	c.Add(
	// 		container.NewVBox(
	// 			widget.NewLabelWithData(lbl),
	// 			slider,
	// 		),
	// 	)
	// 	return c
	// }

	return c
}

func parseBody(body string, links []models.Link, ctx *AppContext) *widget.RichText {
	segments := strings.Split(body, conf.LinkBodyPH)
	bodyRichText := widget.NewRichText()
	for i, t := range segments {
		bodyRichText.Segments = append(bodyRichText.Segments, widgets.NewTSegment(t))
		if len(links) > i {
			link := links[i]
			hL := widgets.NewHyperlinkSegment(link.Label, func() {
				if link.Id != nil {
					ctx.PushWithParam(RouteFromString(link.Route), *link.Id)
				} else {
					ctx.Push(RouteFromString(link.Route))
				}
			})
			bodyRichText.Segments = append(bodyRichText.Segments, hL)
			bodyRichText.Segments = append(bodyRichText.Segments, widgets.NewTSegment(""))
		}
	}

	return bodyRichText
}
