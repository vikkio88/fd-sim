package generators_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamBuilder(t *testing.T) {
	tg := generators.NewTeamGen(generatorsTestSeed)
	team := tg.Team(enums.IT)
	assert.IsType(t, models.Team{}, team)
	assert.NotNil(t, team.Roster)
	assert.Greater(t, team.Roster.Len(), 7)

	assert.NotNil(t, team.Coach)
}
