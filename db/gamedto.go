package db

import (
	"fdsim/models"
	"fdsim/utils"
	"time"

	"gorm.io/gorm"
)

type GameDto struct {
	Id       string `gorm:"primarykey;size:16"`
	SaveName string
	Name     string
	Surname  string
	Age      int
	Fame     int

	StartDate time.Time
	Date      time.Time

	Wage       int64
	YContract  uint8
	Board      int
	Supporters int

	TeamID   *string
	LeagueID *string
	Team     *TeamDto   `gorm:"foreignKey:team_id"`
	League   *LeagueDto `gorm:"foreignKey:league_id"`
}

func DtoFromGame(game *models.Game) GameDto {
	g := GameDto{
		Id:        game.Id,
		SaveName:  game.SaveName,
		Name:      game.Name,
		Surname:   game.Surname,
		Age:       game.Age,
		Fame:      game.Fame.Val(),
		Date:      game.Date,
		StartDate: game.StartDate,
		LeagueID:  &game.LeagueId,
	}

	if game.Team != nil {
		g.TeamID = &game.Team.Id
		g.Wage = game.Wage.Val
		g.YContract = game.YContract
		g.Board = game.Board.Val()
		g.Supporters = game.Supporters.Val()
	}

	return g
}

func (g *GameDto) Game() *models.Game {
	game := models.NewGameWithId(
		g.Id, g.SaveName,
		g.Name, g.Surname, g.Age,
	)
	game.Fame = utils.NewPerc(g.Fame)
	game.LeagueId = *g.LeagueID
	game.Date = g.Date
	game.StartDate = g.StartDate

	if g.Team != nil {
		teamPh := g.Team.Team().PH()
		game.Team = &teamPh
		game.Wage = toMoney(g.Wage)
		game.YContract = g.YContract
		game.Board = utils.NewPerc(g.Board)
		game.Supporters = utils.NewPerc(g.Supporters)
	}

	return game
}
func DtoFromGameWithLeague(game *models.Game, leagueId string) GameDto {
	g := DtoFromGame(game)
	g.LeagueID = &leagueId
	return g
}

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
