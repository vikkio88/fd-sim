package vm

import (
	"fdsim/models"

	"fyne.io/fyne/v2/data/binding"
)

func TeamFromDi(di binding.DataItem) *models.Team {
	v, _ := di.(binding.Untyped).Get()
	return v.(*models.Team)
}
