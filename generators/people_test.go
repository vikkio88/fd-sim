package generators_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeopleGenBuilderPlayerGenerator(t *testing.T) {
	g := generators.NewPeopleGen(generatorsTestSeed)
	pl := g.Player(enums.IT)
	assert.Greater(t, 101, pl.Skill.Val())
	assert.Greater(t, 41, pl.Age)
	assert.Equal(t, "Italian", pl.Country.Nationality())

	pl2 := g.PlayerWithRole(enums.IT, models.GK)
	assert.Equal(t, models.GK, pl2.Role)
}

func TestPeopleGenBuilderCoachGenerator(t *testing.T) {
	g := generators.NewPeopleGen(generatorsTestSeed)
	c := g.Coach(enums.IT)
	assert.Greater(t, 101, c.Skill.Val())
	assert.Greater(t, 81, c.Age)
	assert.Equal(t, "Italian", c.Country.Nationality())
}
