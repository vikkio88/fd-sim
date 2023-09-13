package models_test

import (
	"fdsim/models"
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarketCheckDatesCalculations(t *testing.T) {
	may := utils.NewDate(2023, 5, 15)
	mc := models.CalculateTransferWindowDates(may)
	assert.Equal(t, false, mc.IsOpen())

	midJuly := utils.NewDate(2023, 7, 15)
	mc = models.CalculateTransferWindowDates(midJuly)
	assert.Equal(t, true, mc.IsOpen())

	midDec := utils.NewDate(2023, 12, 15)
	mc = models.CalculateTransferWindowDates(midDec)
	assert.Equal(t, false, mc.IsOpen())

	// mc.ClosestDate()
}
