package services

import (
	"encoding/json"
	"fdsim/conf"
	"fdsim/data"
	"fdsim/db"
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
)

func parsePNPH(dbe db.DbEventDto) []*models.PNPH {
	var result []*models.PNPH
	json.Unmarshal([]byte(dbe.Payload), &result)
	return result
}

func dbEvYoungPlayers(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	youngs := parsePNPH(dbe)

	links := make([]models.Link, len(youngs))
	body := "The following players joined your team from youth squad:"
	for i, rp := range youngs {
		links[i] = models.NewLink(rp.String(), enums.PlayerDetails, &rp.Id)
		body += fmt.Sprintf(" %s", conf.LinkBodyPH)
	}

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(params.TeamName, "hr"),
		fmt.Sprintf("%d players promoted to first team.", len(youngs)),
		body,
		dbe.TriggerDate,
		links,
	)
	return event
}

func dbEvRetiredPlayers(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	retired := parsePNPH(dbe)

	links := make([]models.Link, len(retired))
	body := "The following players retired this year"
	for i, rp := range retired {
		links[i] = models.NewLink(rp.String(), enums.PlayerDetails, &rp.Id)
		body += fmt.Sprintf(" %s", conf.LinkBodyPH)
	}

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(params.TeamName, "hr"),
		fmt.Sprintf("Goodbye! %d Players retired.", len(retired)),
		body,
		dbe.TriggerDate,
		links,
	)
	return event
}

func dbEvSkillChangePlayers(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	var ps [2][]*models.PNPHVals
	json.Unmarshal([]byte(dbe.Payload), &ps)
	improved := ps[0]
	worsened := ps[1]

	links := []models.Link{}

	body := "Few players this year showcased a significant change in their skills."
	if len(improved) > 0 {
		body += "\nThe following players improved quite a lot:"
		for _, v := range improved {
			links = append(links, models.NewLink(fmt.Sprintf("%s : +%d%%", v.String(), v.ValueI), enums.PlayerDetails, &v.Id))
			body += fmt.Sprintf(" %s", conf.LinkBodyPH)
		}
	}
	if len(worsened) > 0 {
		body += "\nThose players skills appear to be worse:"
		for _, v := range worsened {
			links = append(links, models.NewLink(fmt.Sprintf("%s : -%d%%", v.String(), v.ValueI), enums.PlayerDetails, &v.Id))
			body += fmt.Sprintf(" %s", conf.LinkBodyPH)
		}
	}

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(params.TeamName, "training"),
		"Report of your player skill change.",
		body,
		dbe.TriggerDate,
		links,
	)

	return event
}

func dbEvFreeAgent(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	freeagents := parsePNPH(dbe)

	links := make([]models.Link, len(freeagents))
	body := "The following players contracts were not renewed and they are now free agents:"
	for i, rp := range freeagents {
		links[i] = models.NewLink(rp.String(), enums.PlayerDetails, &rp.Id)
		body += fmt.Sprintf(" %s", conf.LinkBodyPH)
	}

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(params.TeamName, "hr"),
		fmt.Sprintf("Contracts expired for %d players.", len(freeagents)),
		body,
		dbe.TriggerDate,
		links,
	)
	return event
}

func dbEvIndividualAwards(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	var statRows []*models.StatRow
	json.Unmarshal([]byte(dbe.Payload), &statRows)

	mvp := statRows[0]
	topScorer := statRows[1]

	links := []models.Link{
		models.NewLink(mvp.Player.String(), enums.PlayerDetails, &mvp.Player.Id),
		models.NewLink(mvp.Team.Name, enums.TeamDetails, &mvp.TeamId),
		models.NewLink(topScorer.Player.String(), enums.PlayerDetails, &topScorer.Player.Id),
		models.NewLink(topScorer.Team.Name, enums.TeamDetails, &topScorer.TeamId),
	}
	body := fmt.Sprintf(`The individual Awards for %s were assigned.
	
	
	MVP
	%s %s
	With an average score of %2.f in %d matches.
	
	
	Top Scorer
	%s %s
	With %d goals in %d matches.`,
		params.LeagueName,
		conf.LinkBodyPH, //mvp link
		conf.LinkBodyPH, // mvp team
		mvp.Score/float64(mvp.Played),
		mvp.Played,
		conf.LinkBodyPH, // scorer link
		conf.LinkBodyPH, // scorer team
		topScorer.Goals, topScorer.Played,
	)

	event.TriggerNews = models.NewNews(
		fmt.Sprintf("%s : Individual Awards Assigned!", params.LeagueName),
		data.GetNewspaper(dbe.Country),
		body,
		dbe.TriggerDate,
		links,
	)
	return event
}

func dbEvTeamAcceptedOffer(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	teamName := params.TeamName
	playerId := params.PlayerId
	playerName := params.PlayerName
	offer := params.ValueF
	offerM := utils.NewEurosFromF(offer)

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(teamName, "hr"),
		fmt.Sprintf("Your offer for %s was accepted.", playerName),
		fmt.Sprintf(`Your offer of %s for the player:
		 %s
was accepted by our board.
You are not allowed to offer him a contract.`, offerM.StringKMB(), conf.LinkBodyPH),
		dbe.TriggerDate,
		[]models.Link{
			playerLink(playerName, playerId),
		},
	)

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		// TODO: move offer status to can offer contract
		of, _ := db.MarketR().GetOffersByPlayerTeamId(playerId, params.FdTeamId)
		of.TeamAccepted = true
		db.MarketR().SaveOffer(of)
	}

	return event
}

func dbEvTeamRefusedOffer(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	teamName := params.TeamName
	playerId := params.PlayerId
	playerName := params.PlayerName
	offer := params.ValueF
	offerM := utils.NewEurosFromF(offer)

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(teamName, "hr"),
		fmt.Sprintf("Your offer for %s was rejected.", playerName),
		fmt.Sprintf(`Your offer of %s for the player:
		 %s
was rejected.`, offerM.StringKMB(), conf.LinkBodyPH),
		dbe.TriggerDate,
		[]models.Link{
			playerLink(playerName, playerId),
		},
	)

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		of, _ := db.MarketR().GetOffersByPlayerTeamId(playerId, params.FdTeamId)
		db.MarketR().DeleteOffer(of)
	}

	return event
}

func dbEvPlayerRefusedContract(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	playerId := params.PlayerId
	playerName := params.PlayerName
	wageOffer := params.ValueF
	wageOfferM := utils.NewEurosFromF(wageOffer)

	event.TriggerEmail = models.NewEmail(
		"agent@footballers.com",
		fmt.Sprintf("%s rejected your contract offer.", playerName),
		fmt.Sprintf(`Your offer of %s for the player:
		 %s
was rejected.`, wageOfferM.StringKMB(), conf.LinkBodyPH),
		dbe.TriggerDate,
		[]models.Link{
			playerLink(playerName, playerId),
		},
	)

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		of, _ := db.MarketR().GetOffersByPlayerTeamId(playerId, params.FdTeamId)
		db.MarketR().DeleteOffer(of)
	}

	return event
}
