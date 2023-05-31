package models_test

import (
	"fdsim/generators"
	"fdsim/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableBuilder(t *testing.T) {
	teams := generators.NewTeamGen(0).Teams(6)
	table := models.NewTable(teams)
	assert.IsType(t, models.Table{}, *table)

	assert.Equal(t, 6, len(table.Rows()))
}
