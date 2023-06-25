package ui

import (
	"fdsim/conf"
	"fdsim/models"
	"fdsim/widgets"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
		Top(NewFborder().Left(backButton(ctx)).Get(centered(h1(title)))).
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
			widget.NewRichText(&widget.TextSegment{Text: news.Body}),
		),
	)
}

func makeEmail(id string, ctx *AppContext) fyne.CanvasObject {
	email := ctx.Db.GameR().GetEmailById(id)
	emailBody := parseBody(email.Body, email.Links, ctx)
	return container.NewMax(
		widget.NewCard(
			email.Subject,
			fmt.Sprintf("%s - %s", email.Date.Format(conf.DateFormatShort), email.Sender),
			emailBody,
		),
	)
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
