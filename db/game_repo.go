package db

import (
	"fdsim/models"

	"gorm.io/gorm"
)

type GameRepo struct {
	g *gorm.DB
}

func NewGameRepo(g *gorm.DB) *GameRepo {
	return &GameRepo{
		g,
	}
}

func (repo *GameRepo) Truncate() {
	repo.g.Where("1 = 1").Delete(&GameDto{})
}

func (repo *GameRepo) All() []*models.Game {
	var dtos []GameDto
	repo.g.Model(&GameDto{}).Find(&dtos)

	ps := make([]*models.Game, len(dtos))
	for i, t := range dtos {
		ps[i] = t.Game()
	}
	return ps
}

func (repo *GameRepo) ById(id string) *models.Game {
	var dto GameDto

	repo.g.Model(&GameDto{}).Find(&dto, "Id = ?", id)

	return dto.Game()
}

func (repo *GameRepo) Create(game *models.Game) {
	dto := DtoFromGame(game)
	repo.g.Create(&dto)
}

func (repo *GameRepo) Update(game *models.Game) {
	dto := DtoFromGame(game)
	repo.g.Save(&dto)
}
