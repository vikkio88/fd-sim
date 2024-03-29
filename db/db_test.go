package db_test

import (
	d "fdsim/db"
	"fdsim/generators"
	"fdsim/libs"
	"fdsim/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLeagueParity(t *testing.T) {
	t.Skip("Slow")
	db := d.NewDb("test.db")
	db.TruncateAll()
	ts := generators.NewTeamGen(time.Now().Unix()).Teams(10)
	db.TeamR().Insert(ts)

	l := models.NewLeague(ts, time.Now())
	db.LeagueR().InsertOne(l)

	dbRestoreLeague, _ := db.LeagueR().ByIdFull(l.Id)
	for _, team := range dbRestoreLeague.Teams() {
		assert.NotNil(t, team.Coach)
		assert.NotEmpty(t, team.Roster.Players())
	}
	assert.Equal(t, l.Name, dbRestoreLeague.Name)
	assert.Equal(t, l.RPointer, dbRestoreLeague.RPointer)
	assert.Equal(t, len(l.Rounds), len(dbRestoreLeague.Rounds))

	r, _ := l.NextRound()
	rFromDb, ok := dbRestoreLeague.NextRound()
	assert.True(t, ok)
	assert.Equal(t, r.Index, rFromDb.Index)
	assert.Equal(t, r.Id, rFromDb.Id)

	rng := libs.NewRng(time.Now().Unix())
	rng1 := libs.NewRng(time.Now().Unix())
	r.Simulate(rng)
	rFromDb.Simulate(rng1)

	l.Update(r)
	dbRestoreLeague.Update(rFromDb)
	assert.Equal(t, l.RPointer, dbRestoreLeague.RPointer)

	db.LeagueR().PostRoundUpdate(rFromDb, dbRestoreLeague)
	stats1 := models.StatsFromRoundResult(rFromDb, dbRestoreLeague.Id)
	db.LeagueR().UpdateStats(stats1)

	//checking if stored stuff is correct
	l, _ = db.LeagueR().ByIdFull(dbRestoreLeague.Id)
	assert.Equal(t, l.RPointer, dbRestoreLeague.RPointer)

	// check if matches are stored correctly
	res := db.LeagueR().RoundByIndex(dbRestoreLeague, rFromDb.Index)
	for _, dbMatchResult := range res.Matches {
		matchInMemory, ok := r.MatchMap[dbMatchResult.Id]
		assert.True(t, ok)
		assert.Equal(t, matchInMemory.Home.Name, dbMatchResult.Home.Name)
		assert.Equal(t, matchInMemory.Away.Name, dbMatchResult.Away.Name)
		result, ok := matchInMemory.Result()
		assert.True(t, ok)
		assert.Equal(t, result.GoalsHome, dbMatchResult.Result.GoalsHome)
		assert.Equal(t, result.GoalsAway, dbMatchResult.Result.GoalsAway)
	}

	// reload League and see if tables Matches
	restoredLeague, _ := db.LeagueR().ByIdFull(dbRestoreLeague.Id)
	assert.Equal(t, restoredLeague.Table.Rows()[0].Team, dbRestoreLeague.Table.Rows()[0].Team)

	r2, ok := restoredLeague.NextRound()
	assert.True(t, ok)
	r2.Simulate(rng1)
	stats2 := models.StatsFromRoundResult(r2, restoredLeague.Id)
	stats1Db := db.LeagueR().GetStats(restoredLeague.Id)
	merged := models.MergeStats(stats1Db, stats2)
	db.LeagueR().UpdateStats(merged)

	// checking if stats update merge works
	statsMergedDb := db.LeagueR().GetStats(restoredLeague.Id)
	for id, value := range merged {
		assert.Equal(t, value.PlayerId, statsMergedDb[id].PlayerId)
		assert.Equal(t, value.Played, statsMergedDb[id].Played)
		assert.Equal(t, value.Goals, statsMergedDb[id].Goals)
	}
}

func TestSingleMatchFetching(t *testing.T) {
	t.Skip("Slow")
	rng := libs.NewRng(time.Now().Unix())
	db := d.NewDb("test.db")

	db.TruncateAll()
	ts := generators.NewTeamGenSeeded(rng).Teams(2)
	db.TeamR().Insert(ts)

	l := models.NewLeague(ts, time.Now())
	db.LeagueR().InsertOne(l)

	league, _ := db.LeagueR().ByIdFull(l.Id)

	r1, ok := league.NextRound()
	assert.True(t, ok)

	m1 := r1.Matches[0]
	m1d := db.LeagueR().GetMatchById(m1.Id)
	assert.Equal(t, m1d.Home.Id, m1.Home.Id)
	assert.Equal(t, m1d.Away.Id, m1.Away.Id)
	assert.Equal(t, m1d.RoundIndex, r1.Index)
	assert.Nil(t, m1d.Result)

	r1.Simulate(rng)
	league.Update(r1)
	db.LeagueR().PostRoundUpdate(r1, league)

	m1afterUpdate := db.LeagueR().GetMatchById(m1.Id)
	assert.Equal(t, m1afterUpdate.Home.Id, m1.Home.Id)
	assert.Equal(t, m1afterUpdate.Away.Id, m1.Away.Id)
	assert.Equal(t, m1afterUpdate.RoundIndex, r1.Index)
	assert.NotNil(t, m1afterUpdate.Result)
}

func TestNonExistingPlayer(t *testing.T) {
	t.Skip("Slow")
	db := d.NewDb("test.db")

	p, ok := db.PlayerR().ById("nonExisting")
	assert.False(t, ok)
	assert.Nil(t, p)
}
