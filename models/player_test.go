package models_test

import (
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildingPlayer(t *testing.T) {
	p := models.NewPlayer("Name", "Surname", 30, models.GK)
	assert.Equal(t, "Name", p.Name)
	assert.Equal(t, "Surname", p.Surname)
	assert.Equal(t, 30, p.Age)
	assert.Equal(t, "Goalkeeper", p.Role.String())
}
