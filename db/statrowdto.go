package db

import "fdsim/models"

type StatRowDto struct {
	PlayerId string `gorm:"primarykey"`
	TeamId   string `gorm:"primarykey;size:16"`
	LeagueId string `gorm:"primarykey;size:16"`
	Played   int
	Goals    int
	Score    float64

	Player *PlayerDto `gorm:"foreignKey:player_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Team   *TeamDto   `gorm:"foreignKey:team_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	League *LeagueDto `gorm:"foreignKey:league_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func DtosFromStatsMap(stats models.StatsMap) []StatRowDto {
	result := []StatRowDto{}
	for _, r := range stats {
		result = append(result, DtoFromStatRow(r))
	}

	return result
}

func DtoFromStatRow(row *models.StatRow) StatRowDto {
	return StatRowDto{
		PlayerId: row.PlayerId,
		TeamId:   row.TeamId,
		LeagueId: row.LeagueId,
		Played:   row.Played,
		Goals:    row.Goals,
		Score:    row.Score,
	}
}

func (row StatRowDto) StatRow() *models.StatRow {
	var player *models.PNPH = nil
	var team *models.TPH = nil

	if row.Player != nil {
		player = row.Player.PlayerPH()
	}

	if row.Team != nil {
		team = row.Team.TeamPH()
	}

	return &models.StatRow{
		PlayerId: row.PlayerId,
		TeamId:   row.TeamId,
		LeagueId: row.LeagueId,
		Played:   row.Played,
		Goals:    row.Goals,
		Score:    row.Score,

		Player: player,
		Team:   team,
	}
}

func StatRowsPhFromDtos(rows []StatRowDto) []*models.StatRowPH {
	result := make([]*models.StatRowPH, len(rows))
	for i, r := range rows {
		result[i] = &models.StatRowPH{
			StatRow: *r.StatRow(),
			Index:   i,
		}
	}

	return result
}

func StatsMapFromDtos(rows []StatRowDto) models.StatsMap {
	result := models.StatsMap{}

	for _, r := range rows {
		result[r.PlayerId] = r.StatRow()
	}

	return result
}
