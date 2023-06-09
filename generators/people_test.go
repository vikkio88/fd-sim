package generators_test

import (
	"fdsim/enums"
	"fdsim/generators"
	"fdsim/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeopleGenBuilderPlayerGenerator(t *testing.T) {
	g := generators.NewPeopleGen(generatorsTestSeed)
	pl := g.Player(enums.IT)
	assert.Greater(t, 101, pl.Skill.Val())
	assert.Less(t, 30, pl.Skill.Val())
	assert.Greater(t, 41, pl.Age)
	assert.Equal(t, "Italian", pl.Country.Nationality())

	assert.Greater(t, pl.Value.Val, int64(1000))
	assert.Greater(t, pl.IdealWage.Val, int64(1000))
	assert.Greater(t, pl.Wage.Val, int64(1000))
	assert.Greater(t, pl.YContract, uint8(0))
	pl2 := g.PlayerWithRole(enums.IT, models.GK)
	assert.Equal(t, models.GK, pl2.Role)
}

func TestPeopleGenBuilderCoachGenerator(t *testing.T) {
	g := generators.NewPeopleGen(generatorsTestSeed)
	c := g.Coach(enums.IT)
	assert.Greater(t, 110, c.Skill.Val())
	assert.Less(t, 30, c.Skill.Val())
	assert.Greater(t, 81, c.Age)
	assert.Equal(t, "Italian", c.Country.Nationality())

	assert.Greater(t, c.IdealWage.Val, int64(1000))
	assert.Greater(t, c.Wage.Val, int64(1000))
	assert.Greater(t, c.YContract, uint8(0))
}

func TestGenerateManyPlayers(t *testing.T) {
	// t.Skip("Long Test")
	g := generators.NewPeopleGen(generatorsTestSeed)
	for i := 0; i < 7000; i++ {
		p := g.Player(enums.IT)
		fmt.Println(p, p.Age, p.Skill, p.Value.StringKMB(), p.IdealWage.StringKMB())
	}
}
