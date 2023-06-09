package ui

import (
	d "fdsim/db"
	vm "fdsim/vm"

	"fyne.io/fyne/v2/data/binding"
)

// TODO:  this is a bit shit but works
var news, emails binding.UntypedList
var dateStr binding.String
var fdTeamId string = ""

// I made them globals to this package as Simulation needs to update the content of the Dashboard

func loadNotifications(db d.IDb) {
	loadEmails(db)
	loadNews(db)
}

func loadNews(db d.IDb) {
	if news == nil {
		news = binding.NewUntypedList()
	} else {
		vm.ClearDataUtList(news)
	}

	newsDb := db.GameR().GetNews()

	for _, n := range newsDb {
		news.Prepend(n)
	}
}

func loadEmails(db d.IDb) {
	if emails == nil {
		emails = binding.NewUntypedList()
	} else {
		vm.ClearDataUtList(emails)
	}
	emailsDb := db.GameR().GetEmails()

	for _, e := range emailsDb {
		emails.Prepend(e)
	}
}

// I hate this, but either I added a bool to all the team and TPH
// or I passed down bools everywhere
func IsFDTeam(teamId string) bool {
	if fdTeamId == "" {
		return false
	}

	return fdTeamId == teamId
}
