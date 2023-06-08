package vm

import (
	"fdsim/models"

	"fyne.io/fyne/v2/data/binding"
)

func PlayerFromDi(di binding.DataItem) *models.Player {
	v, _ := di.(binding.Untyped).Get()
	return v.(*models.Player)
}
