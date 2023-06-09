package models_test

import (
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamBuilder(t *testing.T) {
	team := models.NewTeam("Juventus", "Torino", enums.IT)
	assert.Equal(t, "Juventus", team.Name)
	assert.Equal(t, "Torino", team.City)
	assert.Equal(t, "Italian", team.Country.Nationality())

	assert.IsType(t, utils.Money{}, team.Balance)
	assert.Equal(t, 0.0, team.TransferRatio)
	assert.Nil(t, team.Coach)
	assert.NotNil(t, team.Roster)
}
