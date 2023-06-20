package utils_test

import (
	"fdsim/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatesUtilsLastSunday(t *testing.T) {
	lastSunday := utils.GetLastSunday(2023, time.September)
	assert.Equal(t, time.Sunday, lastSunday.Weekday())
	assert.Equal(t, 24, lastSunday.Day())

	lastSundayMarch2023 := utils.GetLastSunday(2023, time.March)
	assert.Equal(t, time.Sunday, lastSundayMarch2023.Weekday())
	assert.Equal(t, 26, lastSundayMarch2023.Day())
}
func TestDatesUtilsFirstSunday(t *testing.T) {
	firstSunday := utils.GetFirstSunday(2023, time.September)
	assert.Equal(t, time.September, firstSunday.Month())
	assert.Equal(t, time.Sunday, firstSunday.Weekday())
	assert.Equal(t, 3, firstSunday.Day())

	firstSunday2024 := utils.GetFirstSunday(2024, time.January)
	assert.Equal(t, firstSunday2024.Month(), time.January)
	assert.Equal(t, time.Sunday, firstSunday2024.Weekday())
	assert.Equal(t, 7, firstSunday2024.Day())
}
