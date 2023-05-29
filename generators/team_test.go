package generators_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTeamBuilder(t *testing.T) {
	tg := generators.NewTeamGen(generatorsTestSeed)
	team := tg.Team(enums.IT)
	assert.IsType(t, models.Team{}, *team)
	assert.NotNil(t, team.Roster)
	assert.GreaterOrEqual(t, team.Roster.Len(), 17)

	assert.NotNil(t, team.Coach)
}

func TestGeneratingManyTeams(t *testing.T) {
	t.Skip("Long Test")
	tg := generators.NewTeamGen(time.Now().Unix())
	ts := tg.Teams(1000, enums.FR)
	var highest float64 = 0.0
	var lowest float64 = 100.0
	for _, team := range ts {
		if team.Roster.AvgSkill() > highest {
			highest = team.Roster.AvgSkill()
		}

		if team.Roster.AvgSkill() < lowest {
			lowest = team.Roster.AvgSkill()
		}
		fmt.Printf("%s - %d - %.2f\n", team.Name, team.Roster.Len(), team.Roster.AvgSkill())
	}

	fmt.Println(highest, lowest)
}
