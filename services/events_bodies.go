package services

import (
	"fdsim/conf"
	"fdsim/data"
	"fdsim/db"
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
	"time"
)

func seasonOverEvent(params models.EventParams, date time.Time) *Event {
	country := params.Country
	event := NewEvent(date, fmt.Sprintf("Season Over"))

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		db.LeagueR().PostSeason(game)
	}

	event.TriggerNews = models.NewNews(
		"New Season Starting today",
		data.GetNewspaper(country),
		fmt.Sprintf(
			"After an amazing season, today Season %d/%d officially starts.\n%s",
			date.Year(), date.Year()+1, conf.LinkBodyPH,
		),
		date,
		[]models.Link{
			models.NewLink("Previous Season", enums.LeagueHistory, &params.LeagueId),
		},
	)

	return event
}

func leagueFinishedEvent(params models.EventParams, date time.Time) *Event {
	leagueId := params.LeagueId
	leagueName := params.LeagueName
	winnerId := params.TeamId
	winnerName := params.TeamName
	secondId := params.TeamId1
	secondName := params.TeamName1
	thirdId := params.TeamId2
	thirdName := params.TeamName2

	prizeFirst := utils.NewEuros(20_000_000)
	prizeSecond := utils.NewEuros(10_000_000)
	prizeThird := utils.NewEuros(5_000_000)

	event := NewEvent(date, fmt.Sprintf("%s Finished", leagueName))
	title := fmt.Sprintf("%s won %s!", winnerName, leagueName)

	event.TriggerNews = models.NewNews(
		title,
		data.GetNewspaper(params.Country),
		fmt.Sprintf(
			"Today the %s League officially finished, and the winner was %s."+
				"\nFinal Table: %s Winner %s Second Place %s Third Place %s",
			leagueName, winnerName,
			conf.LinkBodyPH, conf.LinkBodyPH,
			conf.LinkBodyPH, conf.LinkBodyPH,
		),
		date,
		[]models.Link{
			models.NewLink(leagueName, enums.League, &leagueId),
			teamLink(winnerName, winnerId),
			teamLink(secondName, secondId),
			teamLink(thirdName, thirdId),
		},
	)

	//TODO: if player team in those 3 top teams notify via email
	if params.IsEmployed {
		position := ""
		subject := ""
		prize := utils.NewEuros(0)
		emailAddr := emailAddrFromTeamName(params.FdTeamName)

		if params.FdTeamId == winnerId {
			position = "first"
			subject = "Amazing job, we won!"
			prize = prizeFirst
		}
		if params.FdTeamId == secondId {
			position = "second"
			subject = "Great stuff, we got to second place!"
			prize = prizeSecond
		}
		if params.FdTeamId == thirdId {
			position = "third"
			subject = "Good job this year, we got to third place."
			prize = prizeThird
		}

		if position != "" {
			event.TriggerEmail = models.NewEmail(
				emailAddr, subject, fmt.Sprintf("We managed to get to %s position. We won %s which was added to our team budget. %s", position, prize.StringKMB(), conf.LinkBodyPH), date,
				[]models.Link{
					models.NewLink(leagueName, enums.League, &leagueId),
				},
			)
		} else {
			event.TriggerEmail = models.NewEmail(
				emailAddr,
				"We did not make it to top 3 this year.",
				fmt.Sprintf("Good job this season, but we did not manage to get to top 3. %s", conf.LinkBodyPH),
				date,
				[]models.Link{
					models.NewLink(leagueName, enums.League, &leagueId),
				},
			)
		}

	}

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		ts := db.TeamR().GetByIds([]string{winnerId, secondId, thirdId})
		for _, t := range ts {
			if t.Id == winnerId {
				t.Balance.Add(prizeFirst)
			}
			if t.Id == secondId {
				t.Balance.Add(prizeSecond)
			}
			if t.Id == thirdId {
				t.Balance.Add(prizeThird)
			}
		}
		for _, t := range ts {
			db.TeamR().Update(t)
		}
	}

	return event
}

func roundPlayedEvent(params models.EventParams, date time.Time) *Event {
	roundIndex := params.Label1
	roundId := params.RoundId
	leagueId := params.LeagueId
	leagueName := params.LeagueName
	desc := fmt.Sprintf("%s - Round %s played", leagueName, roundIndex)
	event := NewEvent(date, desc)

	event.TriggerNews = models.NewNews(
		desc,
		data.GetNewspaper(params.Country),
		fmt.Sprintf(
			"The %sth round of %s  was played today %s, "+
				"Here you can see the updated League table:"+
				" %s Here you can see the round results %s",
			roundIndex,
			leagueName,
			date.Format(conf.DateFormatGame),
			conf.LinkBodyPH,
			conf.LinkBodyPH,
		),
		date,
		[]models.Link{
			models.NewLink(fmt.Sprintf("%s Table", leagueName), enums.League, &leagueId),
			models.NewLink("Round Results", enums.RoundDetails, &roundId),
		},
	)
	return event
}

func contractAccepted(params models.EventParams, date time.Time) *Event {
	teamId := params.TeamId
	teamName := params.TeamName
	ycontract := params.ValueInt
	fdName := params.FdName
	money := utils.NewEurosFromF(params.ValueF)

	title := fmt.Sprintf("%s contract accepted", teamName)

	event := NewEvent(date, title)
	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(teamName),
		fmt.Sprintf("Welcome to %s", teamName),
		fmt.Sprintf(
			"Thanks Mr %s for joining us, we are delighted to have you on board."+
				"Please check our info here: %s CEO of %s",
			fdName,
			conf.LinkBodyPH,
			teamName,
		),
		date,
		[]models.Link{
			teamLink(teamName, teamId),
		},
	)

	event.TriggerNews = models.NewNews(
		fmt.Sprintf("%s hired from %s as Football Director", fdName, teamName),
		data.GetNewspaper(params.Country),
		fmt.Sprintf(
			"A new football director, %s, got hired today from %s. Seems like he signed a %d year(s) contract. %s",
			fdName, teamName, ycontract, conf.LinkBodyPH,
		),
		date,
		[]models.Link{
			teamLink(teamName, teamId),
		},
	)

	event.TriggerFlags = func(f models.Flags) models.Flags {
		f.HasAContractOffer = false
		return f
	}

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		game.SetTeamContract(ycontract, money, &models.TPH{Id: teamId, Name: teamName})
		db.GameR().AddStatRow(models.NewFDStatRow(date, teamId, teamName))
	}

	return event
}

func contractOffered(params models.EventParams, date time.Time) *Event {
	teamId := params.TeamId
	teamName := params.TeamName
	money := utils.NewEurosFromF(params.ValueF)
	years := params.ValueInt

	title := fmt.Sprintf("%s contract offer", teamName)

	event := NewEvent(date, title)
	event.TriggerEmail = models.NewEmailWithAction(
		emailAddrFromTeamName(teamName),
		title,
		fmt.Sprintf(
			"We are willing to offer you %s per year for a lentgth of %d year(s)."+
				"Please consider us for your next job LINK",
			money.StringKMB(),
			years,
		),
		date,
		[]models.Link{
			teamLink(teamName, teamId),
		},
		MakeActionableFromType(
			models.ActionRespondContract,
			date,
			params,
		),
	)
	event.TriggerEmail.SetExpiry(date.Add(2 * enums.A_day))

	event.TriggerFlags = func(f models.Flags) models.Flags {
		f.HasAContractOffer = true
		return f
	}

	return event
}

func transferMarketOpen(params models.EventParams, date time.Time) *Event {
	event := NewEvent(date, "market open")

	session := "Winter"
	// is Summer
	if params.BoolFlag {
		session = "Summer"
	}

	event.TriggerNews = models.NewNews(
		"Transfer Market is open!",
		data.GetNewspaper(params.Country),
		fmt.Sprintf(
			"Today the transfer window for the %s session officially opened!\nIt will stay open until %s",
			session,
			params.Label2,
		),
		date,
		[]models.Link{},
	)

	return event
}
func transferMarketClose(params models.EventParams, date time.Time) *Event {
	event := NewEvent(date, "market closed")

	session := "Winter"
	// is Summer
	if params.BoolFlag {
		session = "Summer"
	}

	event.TriggerNews = models.NewNews(
		"Transfer Market is now closed.",
		data.GetNewspaper(params.Country),
		fmt.Sprintf(
			"Today the transfer market for the %s session is closed!",
			session,
		),
		date,
		[]models.Link{},
	)

	return event
}
