package models_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuleValidation(t *testing.T) {
	module := models.M442
	coach := models.NewCoach("a", "a", 1, enums.IT, models.M442)
	team := generators.NewTeamGen(0).Team(enums.IT)
	team.Coach = &coach

	assert.True(t, module.Validate(team.Lineup().Starting))

	coach2 := models.NewCoach("a", "a", 1, enums.IT, models.M433)
	team.Coach = &coach2
	assert.False(t, module.Validate(team.Lineup().Starting))
}
