package models_test

import (
	"fdsim/enums"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamBuilder(t *testing.T) {
	team := models.NewTeam("Juventus", "Torino", enums.IT)
	assert.Equal(t, "Juventus", team.Name)
	assert.Equal(t, "Torino", team.City)
	assert.Equal(t, "Italian", team.Country.Nationality())
	assert.Nil(t, team.Roster)
	assert.Nil(t, team.Coach)
}
