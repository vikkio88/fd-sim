package utils_test

import (
	"fdsim/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPerc(t *testing.T) {
	p := utils.NewPerc(96)
	assert.Equal(t, 96, p.Val())

	p1 := utils.NewPerc(150)
	assert.Equal(t, 100, p1.Val())

	p2 := utils.NewPerc(-300)
	assert.Equal(t, 0, p2.Val())
}

func TestPercStringer(t *testing.T) {
	p := utils.NewPerc(45)
	assert.Equal(t, "45%", fmt.Sprintf("%s", p))
}
