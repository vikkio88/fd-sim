package db

import (
	"encoding/json"
	"fdsim/models"
	"fdsim/utils"
	"fmt"
	"time"
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

	Flags string
}

func serialiseFlags(f models.Flags) string {
	var result string
	data, _ := json.Marshal(f)
	result = string(data)

	return result
}

func unserialiseFlags(f string) models.Flags {
	var flags models.Flags

	err := json.Unmarshal([]byte(f), &flags)
	if err != nil {
		fmt.Println("Error while getting flags")
		return models.Flags{}
	}

	return flags
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
		Flags:     serialiseFlags(game.Flags),
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
	game.Flags = unserialiseFlags(g.Flags)

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
