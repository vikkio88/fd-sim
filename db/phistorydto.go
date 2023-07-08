package db

import (
	"encoding/json"
	"fdsim/models"
	"fmt"
	"time"
)

type PHistoryDto struct {
	PlayerId string `gorm:"primarykey"`
	Stats    string

	Player PlayerDto
}

func (p *PHistoryDto) Update(stat StatRowDto, gameDate time.Time) {
	existingStats := unserialisePHistoryStats(p.Stats)
	hstat := models.NewPHistoryRow(stat.StatRow(), gameDate)
	existingStats = append(existingStats, newSubRowFromStatRow(hstat))

	p.Stats = serialisePHistoryStats(existingStats)
}

type PHistorySubRow struct {
	LeagueId string
	TeamId   string
	TeamName string
	Played   int
	Goals    int
	Score    float64

	HalfSeason   bool
	TransferCost *string
	StartYear    int
}

func newSubRowFromStatRow(h *models.PHistoryRow) PHistorySubRow {
	return PHistorySubRow{
		LeagueId: h.LeagueId,
		TeamId:   h.TeamId,
		TeamName: h.TeamName,
		Played:   h.Played,
		Goals:    h.Goals,
		Score:    h.Score,

		HalfSeason:   h.HalfSeason,
		TransferCost: h.TransferCost,
		StartYear:    h.StartYear,
	}
}

func newSubRowsFromPHistory(h *models.PHistoryRow) []PHistorySubRow {
	return []PHistorySubRow{
		{
			LeagueId: h.LeagueId,
			TeamId:   h.TeamId,
			TeamName: h.TeamName,
			Played:   h.Played,
			Goals:    h.Goals,
			Score:    h.Score,

			HalfSeason:   h.HalfSeason,
			TransferCost: h.TransferCost,
			StartYear:    h.StartYear,
		},
	}
}

func DtoFromPHistoryRow(h *models.PHistoryRow) PHistoryDto {
	return PHistoryDto{
		PlayerId: h.PlayerId,
		Stats: serialisePHistoryStats(
			newSubRowsFromPHistory(h),
		),
	}
}

func serialisePHistoryStats(pHistorySubRow []PHistorySubRow) string {
	var result string
	data, _ := json.Marshal(pHistorySubRow)
	result = string(data)

	return result
}

func unserialisePHistoryStats(s string) []PHistorySubRow {
	if s == "" {
		return []PHistorySubRow{}
	}

	var result []PHistorySubRow
	data := s
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("error while loading", err)
		return []PHistorySubRow{}
	}

	return result
}
