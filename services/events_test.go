package services_test

import (
	"fdsim/db_test"
	"fdsim/generators"
	"fdsim/models"
	"fdsim/services"
	"fdsim/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEndOfLeague(t *testing.T) {
	tg := generators.NewTeamGen(0)
	ts := tg.Teams(3)
	db := db_test.NewMockDbWithTeams(ts)
	ev := services.LeagueFinished.Event(utils.NewDate(2023, time.May, 20), models.EventParams{
		TeamId:    ts[0].Id,
		TeamName:  "winner",
		TeamId1:   ts[1].Id,
		TeamName1: "second",
		TeamId2:   ts[2].Id,
		TeamName2: "third",
	})

	firstBalance := ts[0].Balance.Value()
	secondBalance := ts[1].Balance.Value()
	thirdBalance := ts[2].Balance.Value()

	game := &models.Game{}
	//TODO: mock db and get info changed
	ev.TriggerChanges(game, db)
	assert.Greater(t, db.Team.Teams[0].Balance.Value(), firstBalance)
	assert.Greater(t, db.Team.Teams[1].Balance.Value(), secondBalance)
	assert.Greater(t, db.Team.Teams[2].Balance.Value(), thirdBalance)
}
