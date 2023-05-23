package models_test

import (
	"fdsim/models"
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildingCoach(t *testing.T) {
	c := models.NewCoach("Name", "Surname", 45, models.M433)
	assert.Equal(t, "Name", c.Name)
	assert.Equal(t, "Surname", c.Surname)
	assert.Equal(t, 45, c.Age)
	assert.Equal(t, "4-3-3", c.Module.String())
}

func TestCoachIsSkillable(t *testing.T) {
	c := models.NewCoach("Name", "Surname", 45, models.M433)

	assert.IsType(t, utils.Perc{}, c.Skill)
	assert.IsType(t, utils.Perc{}, c.Morale)
}
