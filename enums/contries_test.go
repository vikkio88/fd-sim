package enums_test

import (
	"fdsim/enums"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountryToString(t *testing.T) {
	assert.Equal(t, "Italy", enums.IT.String())
	assert.Equal(t, "England", enums.EN.String())
	assert.Equal(t, "France", enums.FR.String())
	assert.Equal(t, "Germany", enums.DE.String())
	assert.Equal(t, "Spain", enums.ES.String())
}

func TestCountryToNationalityString(t *testing.T) {
	assert.Equal(t, "Italian", enums.IT.Nationality())
	assert.Equal(t, "English", enums.EN.Nationality())
	assert.Equal(t, "French", enums.FR.Nationality())
	assert.Equal(t, "German", enums.DE.Nationality())
	assert.Equal(t, "Spanish", enums.ES.Nationality())
}
