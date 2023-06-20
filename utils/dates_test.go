package utils_test

import (
	"fdsim/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatesUtils(t *testing.T) {
	firstSunday := utils.GetFirstSunday(2023, time.September)
	assert.Equal(t, time.September, firstSunday.Month())
	assert.Equal(t, 3, firstSunday.Day())
	lastSunday := utils.GetLastSunday(2023, time.September)
	assert.Equal(t, 24, lastSunday.Day())
	firstSunday2024 := utils.GetFirstSunday(2024, time.January)
	assert.Equal(t, firstSunday2024.Month(), time.January)
	assert.Equal(t, 7, firstSunday2024.Day())
}
