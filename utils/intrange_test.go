package utils_test

import (
	"fdsim/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildIntRange(t *testing.T) {
	r := utils.NewIntRange(1, 3)
	assert.Equal(t, 1, r.Min)
	assert.Equal(t, 3, r.Max)

	r2 := utils.NewIntRange(3, 2)
	assert.Equal(t, 2, r2.Min)
	assert.Equal(t, 3, r2.Max)
}

func TestBuildIntRangeFromString(t *testing.T) {
	r := utils.NewIntRangeFromStr("1..3")
	assert.Equal(t, 1, r.Min)
	assert.Equal(t, 3, r.Max)

	r1 := utils.NewIntRangeFromStr("something")
	assert.Equal(t, 0, r1.Min)
	assert.Equal(t, 0, r1.Max)

	r2 := utils.NewIntRangeFromStr("3..1")
	assert.Equal(t, 1, r2.Min)
	assert.Equal(t, 3, r2.Max)
}
