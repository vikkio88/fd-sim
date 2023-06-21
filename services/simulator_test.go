package services_test

import (
	"fdsim/db_test"
	"fdsim/models"
	"fdsim/services"
	"fdsim/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func NewMockedGame(startDate, date time.Time) *models.Game {
	g := models.NewGameWithId("mockGame", "SomeGame", "Testing", "User", 35)
	g.StartDate = startDate
	g.Date = date

	return g
}

func NewMockDb() *db_test.MockDb {
	return &db_test.MockDb{}
}

func TestSimulatorBuilder(t *testing.T) {
	sim := services.NewSimulator(NewMockedGame(time.Now(), time.Now()), NewMockDb())
	assert.IsType(t, &services.Simulator{}, sim)
}

func TestSimulatorAdvancingTime(t *testing.T) {
	startDate := utils.NewDate(2023, time.July, 1)
	game := NewMockedGame(startDate, startDate)
	sim := services.NewSimulator(game, db_test.NewMockDbSeeded(0))
	days := 1
	sim.Simulate(days)
	assert.True(t, game.StartDate.Equal(startDate))
	assert.Equal(t, 2, game.Date.Day())

	days = 7
	sim.Simulate(days)
	assert.True(t, game.StartDate.Equal(startDate))
	assert.Equal(t, 9, game.Date.Day())
	assert.Equal(t, time.July, game.Date.Month())
}
