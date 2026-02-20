package services

import (
	"encoding/json"
	"fdsim/conf"
	"fdsim/data"
	"fdsim/db"
	"fdsim/enums"
	"fdsim/libs"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
	"strings"
)

func parsePNPH(dbe db.DbEventDto) []*models.PNPH {
	var result []*models.PNPH
	json.Unmarshal([]byte(dbe.Payload), &result)
	return result
}

func dbEvYoungPlayers(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	youngs := parsePNPH(dbe)

	links := make([]models.Link, len(youngs))
	var body strings.Builder
	body.WriteString("The following players joined your team from youth squad:")
	for i, rp := range youngs {
		links[i] = models.NewLink(rp.String(), enums.PlayerDetails, &rp.Id)
		body.WriteString(fmt.Sprintf(" %s", conf.LinkBodyPH))
	}

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(params.TeamName, "hr"),
		fmt.Sprintf("%d players promoted to first team.", len(youngs)),
		body.String(),
		dbe.TriggerDate,
		links,
	)
	return event
}

func dbEvRetiredPlayers(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	retired := parsePNPH(dbe)

	links := make([]models.Link, len(retired))
	var body strings.Builder
	body.WriteString("The following players retired this year")
	for i, rp := range retired {
		links[i] = models.NewLink(rp.String(), enums.PlayerDetails, &rp.Id)
		body.WriteString(fmt.Sprintf(" %s", conf.LinkBodyPH))
	}

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(params.TeamName, "hr"),
		fmt.Sprintf("Goodbye! %d Players retired.", len(retired)),
		body.String(),
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

	var body strings.Builder
	body.WriteString("Few players this year showcased a significant change in their skills.")
	if len(improved) > 0 {
		body.WriteString("\nThe following players improved quite a lot:")
		for _, v := range improved {
			links = append(links, models.NewLink(fmt.Sprintf("%s : +%d%%", v.String(), v.ValueI), enums.PlayerDetails, &v.Id))
			body.WriteString(fmt.Sprintf(" %s", conf.LinkBodyPH))
		}
	}
	if len(worsened) > 0 {
		body.WriteString("\nThose players skills appear to be worse:")
		for _, v := range worsened {
			links = append(links, models.NewLink(fmt.Sprintf("%s : -%d%%", v.String(), v.ValueI), enums.PlayerDetails, &v.Id))
			body.WriteString(fmt.Sprintf(" %s", conf.LinkBodyPH))
		}
	}

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(params.TeamName, "training"),
		"Report of your player skill change.",
		body.String(),
		dbe.TriggerDate,
		links,
	)

	return event
}

func dbEvFreeAgent(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	freeagents := parsePNPH(dbe)

	links := make([]models.Link, len(freeagents))
	var body strings.Builder
	body.WriteString("The following players contracts were not renewed and they are now free agents:")
	for i, rp := range freeagents {
		links[i] = models.NewLink(rp.String(), enums.PlayerDetails, &rp.Id)
		body.WriteString(fmt.Sprintf(" %s", conf.LinkBodyPH))
	}

	event.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(params.TeamName, "hr"),
		fmt.Sprintf("Contracts expired for %d players.", len(freeagents)),
		body.String(),
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
			playerSubLink(playerName, playerId, enums.PlayerDTransferTab),
		},
	)

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		// TODO: move offer status to can offer contract
		of, _ := db.MarketR().GetOffersByPlayerTeamId(playerId, params.FdTeamId)
		of.TeamAccepted = true
		of.LastUpdate = game.Date
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

func dbEvPlayerAcceptedContract(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	playerId := params.PlayerId
	playerName := params.PlayerName
	wageOffer := params.ValueF
	wageOfferM := utils.NewEurosFromF(wageOffer)
	yContract := params.ValueInt

	event.TriggerEmail = models.NewEmail(
		getPlayerEmail(playerName, dbe.Country),
		fmt.Sprintf("%s accepted your contract offer.", playerName),
		fmt.Sprintf(`You offered a wage of %s for %d year(s) for the player:
		 %s
was accepted.`, wageOfferM.StringKMB(), yContract, conf.LinkBodyPH),
		dbe.TriggerDate,
		[]models.Link{
			playerSubLink(playerName, playerId, enums.PlayerDTransferTab),
		},
	)

	event.TriggerChanges = func(game *models.Game, db db.IDb) {
		of, _ := db.MarketR().GetOffersByPlayerTeamId(playerId, params.FdTeamId)
		of.PlayerAccepted = true
		of.LastUpdate = game.Date
		db.MarketR().SaveOffer(of)
	}

	return event
}

func dbEvPlayerRefusedContract(dbe db.DbEventDto, event *Event, params models.EventParams) *Event {
	playerId := params.PlayerId
	playerName := params.PlayerName
	wageOffer := params.ValueF
	wageOfferM := utils.NewEurosFromF(wageOffer)
	yContract := params.ValueInt

	event.TriggerEmail = models.NewEmail(
		getPlayerEmail(playerName, dbe.Country),
		fmt.Sprintf("%s rejected your contract offer.", playerName),
		fmt.Sprintf(`You offered a wage of %s for %d year(s) for the player:
		 %s
was rejected.`, wageOfferM.StringKMB(), yContract, conf.LinkBodyPH),
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

func dbEvTransferHappening(dbe db.DbEventDto, event *Event, params models.EventParams, d db.IDb) *Event {
	ev := NewEmptyEvent()
	var offer *models.Offer
	json.Unmarshal([]byte(dbe.Payload), &offer)

	// Perform transfer
	result := d.MarketR().ApplyTransfer(offer)

	if !result.Success {
		ev.TriggerEmail = models.NewEmail(
			emailAddrFromTeamName(result.Team.Name, "finance"),
			fmt.Sprintf("We blocked the transfer of %s", result.Player.String()),
			fmt.Sprintf(
				`Finance department blocked the transfer of the player:
%s
as the offer of %s would go over the budget allocated for transfer market.`,
				conf.LinkBodyPH, result.Bid.StringKMB()),
			offer.TransferDate,
			[]models.Link{
				playerLink(result.Player.String(), result.Player.Id),
			})
		ev.TriggerChanges = func(game *models.Game, db db.IDb) {
			rng := libs.NewRngAutoSeeded()
			// made board trust you less
			game.Board.SetVal(game.Board.Val() - rng.UInt(2, 5))
			db.GameR().Update(game)
		}
		return ev
	}

	links := []models.Link{
		playerLink(result.Player.String(), result.Player.Id),
	}

	teamBodyPart := "He was previously free agent. So he was signed for free."
	if result.PreviousTeam != nil {
		teamBodyPart = fmt.Sprintf("He previously played for %s, and they accepted an offer of %s.", conf.LinkBodyPH, result.Bid.StringKMB())
		links = append(links, teamLink(result.PreviousTeam.Name, result.PreviousTeam.Id))
	}

	// success
	ev.TriggerEmail = models.NewEmail(
		emailAddrFromTeamName(result.Team.Name, "hr"),
		fmt.Sprintf("Welcome %s to our team.", result.Player.String()),
		fmt.Sprintf(
			`We welcome the player:
%s
to our team.
%s`,
			conf.LinkBodyPH, teamBodyPart),
		offer.TransferDate,
		links,
	)

	ev.TriggerChanges = func(game *models.Game, db db.IDb) {
		rng := libs.NewRngAutoSeeded()

		if rng.Chance(result.PlayerSkill) {
			game.Board.SetVal(game.Board.Val() + rng.UInt(1, 5))
			game.Fame.SetVal(game.Board.Val() + rng.UInt(1, 2))
		}

		//TODO: add to FD History/Stats
	}

	return ev
}
