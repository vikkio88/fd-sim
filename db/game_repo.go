package db

import (
	"fdsim/models"
	"time"

	"gorm.io/gorm"
)

type GameRepo struct {
	g *gorm.DB
}

// TruncateNotifications implements IGameRepo.
func (repo *GameRepo) TruncateNotifications() {
	repo.g.Where("1=1").Delete(&EmailDto{})
	repo.g.Where("1=1").Delete(&NewsDto{})
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
	repo.g.Model(&GameDto{}).Preload(teamRel).Find(&dtos)

	ps := make([]*models.Game, len(dtos))
	for i, t := range dtos {
		ps[i] = t.Game()
	}
	return ps
}

func (repo *GameRepo) ById(id string) *models.Game {
	var dto GameDto

	repo.g.Model(&GameDto{}).Preload(teamRel).Find(&dto, "Id = ?", id)

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

func (repo *GameRepo) AddStatRow(row *models.FDStatRow) {
	dto := DtoFromFDStatRow(row)
	repo.g.Create(&dto)
}

// This is to get Emails with actions due a certain date
func (repo *GameRepo) GetActionsDueByDate(date time.Time) []*models.Idable {
	var dtos []EmailDto
	repo.g.Where("expires = ? AND action is not null and decision is null", date).Find(&dtos)

	result := make([]*models.Idable, len(dtos))
	for i, dto := range dtos {
		result[i] = dto.Idable()
	}

	return result
}

func (repo *GameRepo) AddEmails(emails []*models.Email) {
	if len(emails) < 1 {
		return
	}
	dtos := make([]EmailDto, len(emails))
	for i, e := range emails {
		dtos[i] = DtoFromEmail(e)

	}
	repo.g.Model(&EmailDto{}).Create(&dtos)
}

func (repo *GameRepo) UpdateEmail(email *models.Email) {
	dto := DtoFromEmail(email)
	repo.g.Model(&EmailDto{}).Where("id = ?", dto.Id).Save(&dto)
}

func (repo *GameRepo) GetEmails() []*models.Email {
	var dtos []EmailDto
	repo.g.Model(&EmailDto{}).Order("date desc").Find(&dtos)

	result := make([]*models.Email, len(dtos))
	for i, dto := range dtos {
		result[i] = dto.Email()
	}

	return result
}

func (repo *GameRepo) GetEmailById(id string) *models.Email {
	var dto EmailDto
	repo.g.Model(&EmailDto{}).First(&dto, "id = ?", id)
	return dto.Email()
}

func (repo *GameRepo) MarkEmailAsRead(id string) {
	repo.g.Model(&EmailDto{}).Where("id = ?", id).Update("read", true)
}

func (repo *GameRepo) DeleteEmail(id string) {
	repo.g.Where("id = ?", id).Delete(&EmailDto{})
}

func (repo *GameRepo) AddNews(news []*models.News) {
	if len(news) < 1 {
		return
	}
	dtos := make([]NewsDto, len(news))
	for i, n := range news {
		dtos[i] = DtoFromNews(n)

	}
	repo.g.Model(&NewsDto{}).Create(&dtos)
}

func (repo *GameRepo) GetNews() []*models.News {
	var dtos []NewsDto
	repo.g.Model(&NewsDto{}).Order("date desc").Find(&dtos)

	result := make([]*models.News, len(dtos))
	for i, dto := range dtos {
		result[i] = dto.News()
	}

	return result
}

func (repo *GameRepo) GetNewsById(id string) *models.News {
	var dto NewsDto
	repo.g.Model(&NewsDto{}).First(&dto, "id = ?", id)
	return dto.News()
}

func (repo *GameRepo) MarkNewsAsRead(id string) {
	repo.g.Model(&NewsDto{}).Where("id = ?", id).Update("read", true)
}

func (repo *GameRepo) DeleteNews(id string) {
	repo.g.Where("id = ?", id).Delete(&NewsDto{})
}

func (repo *GameRepo) DeleteAllNews() {
	repo.g.Where("1 = 1").Delete(&NewsDto{})
}

func (repo *GameRepo) GetFDHistory() []*models.FDHistoryRow {
	var dtos []FDHistoryDto

	repo.g.Model(&FDHistoryDto{}).Find(&dtos)
	result := make([]*models.FDHistoryRow, len(dtos))
	for i, r := range dtos {
		result[i] = r.FDHistoryRow()
	}

	return result
}

func (repo *GameRepo) GetFDStats() *models.FDStatRow {
	var stat FDStatRowDto
	repo.g.Model(&FDStatRowDto{}).Order("hired_date desc").First(&stat)

	return stat.FDStatRow()
}

// GetEvents implements IGameRepo.
func (repo *GameRepo) GetEvents(currentDate time.Time) []DbEventDto {
	var ev []DbEventDto
	// get then delete
	repo.g.Model(&DbEventDto{}).Where("trigger_date < ?", currentDate).Find(&ev)
	repo.g.Where("trigger_date < ?", currentDate).Delete(&DbEventDto{})

	return ev
}

// TODO: remove if not used
func (*GameRepo) StoreEvents([]DbEventDto) {
}

func (repo *GameRepo) StoreEvent(ev DbEventDto) {
	repo.g.Model(&DbEventDto{}).Create(&ev)
}
