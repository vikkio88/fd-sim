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

type FDHistoryDto struct {
	Year       int        `gorm:"primarykey"`
	Month      time.Month `gorm:"primarykey"`
	Day        int        `gorm:"primarykey"`
	TeamId     string     `gorm:"primarykey"`
	TeamName   string     `gorm:"primarykey"`
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
		Year:     r.HiredDate.Year(),
		Month:    r.HiredDate.Month(),
		Day:      r.HiredDate.Day(),
		TeamId:   r.TeamId,
		TeamName: r.TeamName,

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

func (h *FDHistoryDto) UpdateEndOfSeason(leagueId, leagueName string, wage utils.Money) {
	h.LeagueId = leagueId
	h.LeagueName = leagueName
	h.Wage = wage.Value()
}
