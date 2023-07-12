package db

import (
	"fdsim/models"
	"fdsim/utils"
	"time"
)

type FDStatRowDto struct {
	TeamId    string
	TeamName  string
	HiredDate time.Time

	PlayersSigned int
	PlayersSold   int
	CoachesSigned int
	CoachesSacked int

	MaxSpent    float64
	TotalSpent  float64
	MaxCashed   float64
	TotalCashed float64
}

func DtoFromFDStatRow(r *models.FDStatRow) FDStatRowDto {
	return FDStatRowDto{
		TeamId:    r.TeamId,
		TeamName:  r.TeamName,
		HiredDate: r.HiredDate,
	}
}

func (r *FDStatRowDto) FDStatRow() *models.FDStatRow {
	return &models.FDStatRow{
		TeamId:    r.TeamId,
		TeamName:  r.TeamName,
		HiredDate: r.HiredDate,

		PlayersSigned: r.PlayersSigned,
		PlayersSold:   r.PlayersSold,
		CoachesSigned: r.CoachesSigned,
		CoachesSacked: r.CoachesSacked,

		MaxSpent:    utils.NewEurosFromF(r.MaxSpent),
		TotalSpent:  utils.NewEurosFromF(r.TotalSpent),
		MaxCashed:   utils.NewEurosFromF(r.MaxCashed),
		TotalCashed: utils.NewEurosFromF(r.TotalCashed),
	}

}

type FDHistoryDto struct {
	Year       int        `gorm:"primarykey"`
	Month      time.Month `gorm:"primarykey"`
	Day        int        `gorm:"primarykey"`
	StartDate  time.Time
	EndDate    time.Time
	TeamId     string `gorm:"primarykey"`
	TeamName   string `gorm:"primarykey"`
	LeagueId   string
	LeagueName string
	Wage       float64

	PlayersSigned int
	PlayersSold   int
	CoachesSigned int
	CoachesSacked int

	MaxSpent    float64
	TotalSpent  float64
	MaxCashed   float64
	TotalCashed float64
}

func NewFDHistoryDto(r FDStatRowDto) FDHistoryDto {
	return FDHistoryDto{
		Year:      r.HiredDate.Year(),
		Month:     r.HiredDate.Month(),
		Day:       r.HiredDate.Day(),
		StartDate: r.HiredDate,
		TeamId:    r.TeamId,
		TeamName:  r.TeamName,

		PlayersSigned: r.PlayersSigned,
		PlayersSold:   r.PlayersSold,
		CoachesSigned: r.CoachesSigned,
		CoachesSacked: r.CoachesSacked,

		MaxSpent:    r.MaxSpent,
		TotalSpent:  r.TotalSpent,
		MaxCashed:   r.MaxCashed,
		TotalCashed: r.TotalCashed,
	}
}

func (h *FDHistoryDto) UpdateEndOfSeason(leagueId, leagueName string, wage utils.Money, gameDate time.Time) {
	h.LeagueId = leagueId
	h.LeagueName = leagueName
	h.Wage = wage.Value()
	h.EndDate = gameDate
}

func (h *FDHistoryDto) FDHistoryRow() *models.FDHistoryRow {
	return &models.FDHistoryRow{
		StartDate:  h.StartDate,
		EndDate:    h.EndDate,
		TeamId:     h.TeamId,
		TeamName:   h.TeamName,
		LeagueId:   h.LeagueId,
		LeagueName: h.LeagueName,
		Wage:       utils.NewEurosFromF(h.Wage),

		PlayersSigned: h.PlayersSigned,
		PlayersSold:   h.PlayersSold,
		CoachesSigned: h.CoachesSigned,
		CoachesSacked: h.CoachesSacked,

		MaxSpent:    utils.NewEurosFromF(h.MaxSpent),
		TotalSpent:  utils.NewEurosFromF(h.TotalSpent),
		MaxCashed:   utils.NewEurosFromF(h.MaxCashed),
		TotalCashed: utils.NewEurosFromF(h.TotalCashed),
	}
}
