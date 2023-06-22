package vm

import (
	"fdsim/models"

	"fyne.io/fyne/v2/data/binding"
)

func NewsFromDi(di binding.DataItem) *models.News {
	v, _ := di.(binding.Untyped).Get()
	return v.(*models.News)
}

func EmailFromDi(di binding.DataItem) *models.Email {
	v, _ := di.(binding.Untyped).Get()
	return v.(*models.Email)
}
