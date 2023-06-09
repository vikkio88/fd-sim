package models_test

import (
	"fdsim/enums"
	"fdsim/models"
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildingCoach(t *testing.T) {
	c := models.NewCoach("Name", "Surname", 45, enums.ES, models.M433)
	assert.Equal(t, "Name", c.Name)
	assert.Equal(t, "Surname", c.Surname)
	assert.Equal(t, "Spanish", c.Country.Nationality())
	assert.Equal(t, 45, c.Age)
	assert.Equal(t, "4-3-3", c.Module.String())

	// idable
	assert.Contains(t, c.Id, "cmId_")

	// Wage
	assert.Equal(t, c.IdealWage.Val, utils.Money{}.Val)
}

func TestCoachIsSkillable(t *testing.T) {
	c := models.NewCoach("Name", "Surname", 45, enums.DE, models.M433)

	assert.IsType(t, utils.Perc{}, c.Skill)
	assert.IsType(t, utils.Perc{}, c.Morale)

	c.SetVals(65, 70, 30)

	assert.Equal(t, 65, c.Skill.Val())
	assert.Equal(t, 70, c.Morale.Val())
	assert.Equal(t, 30, c.Fame.Val())
}
