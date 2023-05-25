package models_test

import (
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineupBuilder(t *testing.T) {
	l := models.NewLineup(models.M442, models.NewRolePPHMap())
	assert.IsType(t, models.Lineup{}, l)
}
