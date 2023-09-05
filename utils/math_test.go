package utils_test

import (
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathExp(t *testing.T) {
	value := utils.ExpFactor(30, 60, 50)
	assert.GreaterOrEqual(t, value, 20.0)
	value = utils.ExpFactor(40, 70, 100)
	assert.GreaterOrEqual(t, value, 70.0)
	value = utils.ExpFactor(50, 70, 100)
	assert.GreaterOrEqual(t, value, 80.0)
	value = utils.ExpFactor(60, 70, 100)
	assert.GreaterOrEqual(t, value, 90.0)
	value = utils.ExpFactor(70, 70, 100)
	assert.Equal(t, value, 100.0)
	value = utils.ExpFactor(80, 70, 100)
	assert.GreaterOrEqual(t, value, 100.0)
}
