package ui

import (
	d "fdsim/db"
	"fdsim/models"
	vm "fdsim/vm"
	"fdsim/widgets"
	"fmt"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func makeNotificationsTabs(ctx *AppContext, navigate func(route AppRoute, param any)) *container.AppTabs {
	newsMailsTabs := container.NewAppTabs(
		container.NewTabItemWithIcon("News", widgets.Icon("newspaper").Resource, makeNewsTab(news, ctx.Db, navigate)),
		container.NewTabItemWithIcon("Emails", theme.MailComposeIcon(), makeEmailsTab(emails, ctx.Db, navigate)),
	)
	emails.AddListener(binding.NewDataListener(func() {
		e, _ := emails.Get()
		unread := countUnread(e, true)
		if unread > 0 {
			newsMailsTabs.Items[1].Text = fmt.Sprintf("Email (%d)", unread)
		}
		newsMailsTabs.Refresh()
	}))
	news.AddListener(binding.NewDataListener(func() {
		n, _ := news.Get()
		unread := countUnread(n, false)
		if unread > 0 {
			newsMailsTabs.Items[0].Text = fmt.Sprintf("News (%d)", unread)
			newsMailsTabs.Refresh()
		}
	}))

	// This will remove the (1) sign on the text of the tab when you go to it
	newsMailsTabs.OnSelected = func(ti *container.TabItem) {
		re := regexp.MustCompile(`\s\(\d+\)$`)
		ti.Text = re.ReplaceAllString(ti.Text, "")
	}
	return newsMailsTabs
}

// Functions that will count Unred Notifications
func countUnread(notifications []any, isEmail bool) int {
	c := 0
	for _, e := range notifications {
		//TODO: can change this to the golang type inference
		if isEmail && !e.(*models.Email).Read {
			c++
			continue
		}

		if !isEmail && !e.(*models.News).Read {
			c++
		}
	}

	return c
}

func makeNewsTab(news binding.UntypedList, db d.IDb, navigate NavigateWithParamFunc) fyne.CanvasObject {
	list := widget.NewListWithData(
		news,
		func() fyne.CanvasObject {
			return container.NewMax(
				NewFborder().
					Right(
						widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
					).
					Get(
						container.NewVBox(
							widget.NewLabel(""),
							// widget.NewLabel(""),
						),
					),
			)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			newsI := vm.NewsFromDi(di)
			newsContainer := co.(*fyne.Container).Objects[0].(*fyne.Container)
			newsInfoCtr := newsContainer.Objects[0].(*fyne.Container) //.Objects[0].(*fyne.Container)
			mainLbl := newsInfoCtr.Objects[0].(*widget.Label)
			mainLbl.SetText(newsI.String())
			mainLbl.TextStyle = fyne.TextStyle{Bold: !newsI.Read}
			mainLbl.Refresh()
			deleteBtn := newsContainer.Objects[1].(*widget.Button)
			deleteBtn.OnTapped = func() {
				db.GameR().DeleteNews(newsI.Id)
				items, _ := news.Get()
				index := slices.IndexFunc(items, func(item any) bool {
					e := item.(*models.News)
					return e.Id == newsI.Id
				})
				items = append(items[:index], items[index+1:]...)
				news.Set(items)
			}
		})

	list.OnSelected = func(id widget.ListItemID) {
		di, _ := news.GetItem(id)
		news := vm.NewsFromDi(di)
		if !news.Read {
			news.Read = true
			list.Refresh()
			db.GameR().MarkNewsAsRead(news.Id)
		}
		list.UnselectAll()
		navigate(News, news.Id)
	}
	return NewFborder().
		Top(
			container.NewPadded(
				leftAligned(widget.NewButtonWithIcon("All", theme.DeleteIcon(), func() {
					vm.ClearDataUtList(news)
					db.GameR().DeleteAllNews()
				})),
			),
		).
		Get(
			list,
		)
}

func makeEmailsTab(emails binding.UntypedList, db d.IDb, navigate NavigateWithParamFunc) fyne.CanvasObject {
	list := widget.NewListWithData(
		emails,
		func() fyne.CanvasObject {
			return container.NewMax(
				NewFborder().
					Left(widget.NewIcon(theme.MailComposeIcon())).
					Right(
						container.NewHBox(
							widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
							widget.NewIcon(theme.InfoIcon()),
						),
					).
					Get(
						container.NewVBox(
							widget.NewLabel(""),
						),
					),
			)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			email := vm.EmailFromDi(di)
			emailCtr := co.(*fyne.Container).Objects[0].(*fyne.Container)
			mailInfoCtr := emailCtr.Objects[0].(*fyne.Container) //.Objects[0].(*fyne.Container)
			mainLbl := mailInfoCtr.Objects[0].(*widget.Label)
			mainLbl.SetText(email.String())
			mainLbl.TextStyle = fyne.TextStyle{Bold: !email.Read}
			mainLbl.Refresh()

			leftIcon := emailCtr.Objects[1].(*widget.Icon)
			if email.Read {
				leftIcon.SetResource(widgets.Icon("email_read").Resource)
			}
			deleteBtn := emailCtr.Objects[2].(*fyne.Container).Objects[0].(*widget.Button)
			infoIcon := emailCtr.Objects[2].(*fyne.Container).Objects[1].(*widget.Icon)
			infoIcon.Hide()
			deleteBtn.OnTapped = func() {
				db.GameR().DeleteEmail(email.Id)
				items, _ := emails.Get()
				index := slices.IndexFunc(items, func(item any) bool {
					e := item.(*models.Email)
					return e.Id == email.Id
				})
				items = append(items[:index], items[index+1:]...)
				emails.Set(items)
			}

			if _, actionable := email.IsActionable(); actionable {
				infoIcon.Show()
				deleteBtn.Hide()
			} else {
				deleteBtn.Show()
			}
		})

	list.OnSelected = func(id widget.ListItemID) {
		di, _ := emails.GetItem(id)
		email := vm.EmailFromDi(di)
		if !email.Read {
			email.Read = true
			list.Refresh()
			db.GameR().MarkEmailAsRead(email.Id)
		}
		list.UnselectAll()
		navigate(Email, email.Id)
	}
	return list
}
