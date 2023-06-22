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

func (repo *GameRepo) AddEmails(emails []*models.Email) {
	if len(emails) < 1 {
		return
	}
	dtos := make([]EmailDto, len(emails))
	for i, e := range emails {
		dtos[i] = DtoFromEmail(e)

	}
	repo.g.Model(&EmailDto{}).Create(dtos)
}

func (repo *GameRepo) GetEmails() []*models.Email {
	return []*models.Email{}
}

func (*GameRepo) GetEmailById(id string) *models.Email {
	return nil
}

func (*GameRepo) MarkEmailAsRead(id string) {

}

func (*GameRepo) DeleteEmail(id string) {

}

func (repo *GameRepo) AddNews(news []*models.News) {
	if len(news) < 1 {
		return
	}
	dtos := make([]NewsDto, len(news))
	for i, n := range news {
		dtos[i] = DtoFromNews(n)

	}
	repo.g.Model(&NewsDto{}).Create(dtos)
}

func (*GameRepo) GetNews() []*models.News {
	return []*models.News{}
}

func (*GameRepo) GetNewsById(id string) *models.News {
	return nil
}

func (*GameRepo) MarkNewsAsRead(id string) {

}

func (*GameRepo) DeleteNews(id string) {

}
