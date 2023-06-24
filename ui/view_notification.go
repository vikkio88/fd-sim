package ui

import (
	"fdsim/conf"
	"fmt"

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
	return container.NewMax(
		widget.NewCard(
			email.Subject,
			fmt.Sprintf("%s - %s", email.Date.Format(conf.DateFormatShort), email.Sender),
			widget.NewRichText(&widget.TextSegment{Text: email.Body}),
		),
	)
}
