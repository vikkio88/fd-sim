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

func (p *PHistoryDto) HistoryRows() []*models.PHistoryRow {
	if p.Stats == "" {
		return []*models.PHistoryRow{}
	}

	stats := unserialisePHistoryStats(p.Stats)
	if len(stats) < 1 {
		return []*models.PHistoryRow{}
	}

	result := make([]*models.PHistoryRow, len(stats))
	for i, s := range stats {
		result[i] = s.PHistoryRow()
	}

	return result
}

func (p *PHistoryDto) Update(stat StatRowDto, leagueName string, gameDate time.Time) {
	existingStats := unserialisePHistoryStats(p.Stats)
	hstat := models.NewPHistoryRow(stat.StatRow(), leagueName, gameDate)
	existingStats = append(existingStats, newSubRowFromStatRow(hstat))

	p.Stats = serialisePHistoryStats(existingStats)
}

type PHistorySubRow struct {
	LeagueId   string
	LeagueName string
	TeamId     string
	TeamName   string
	Played     int
	Goals      int
	Score      float64

	HalfSeason   bool
	TransferCost *string
	StartYear    int
}

func newSubRowFromStatRow(h *models.PHistoryRow) PHistorySubRow {
	return PHistorySubRow{
		LeagueId:   h.LeagueId,
		LeagueName: h.LeagueName,
		TeamId:     h.TeamId,
		TeamName:   h.TeamName,
		Played:     h.Played,
		Goals:      h.Goals,
		Score:      h.Score,

		HalfSeason:   h.HalfSeason,
		TransferCost: h.TransferCost,
		StartYear:    h.StartYear,
	}
}

func (h *PHistorySubRow) PHistoryRow() *models.PHistoryRow {
	return &models.PHistoryRow{
		LeagueId:     h.LeagueId,
		LeagueName:   h.LeagueName,
		TeamId:       h.TeamId,
		TeamName:     h.TeamName,
		Played:       h.Played,
		Goals:        h.Goals,
		Score:        h.Score,
		HalfSeason:   h.HalfSeason,
		TransferCost: h.TransferCost,
		StartYear:    h.StartYear,
	}
}
func newSubRowsFromPHistory(h *models.PHistoryRow) []PHistorySubRow {
	return []PHistorySubRow{
		{
			LeagueId:   h.LeagueId,
			LeagueName: h.LeagueName,
			TeamId:     h.TeamId,
			TeamName:   h.TeamName,
			Played:     h.Played,
			Goals:      h.Goals,
			Score:      h.Score,

			HalfSeason:   h.HalfSeason,
			TransferCost: h.TransferCost,
			StartYear:    h.StartYear,
		},
	}
}

func NewEmptyHistorySubRow(leagueId, leagueName, teamId, teamName string, year int) []PHistorySubRow {
	return []PHistorySubRow{
		{
			LeagueId:   leagueId,
			LeagueName: leagueName,
			TeamId:     teamId,
			TeamName:   teamName,
			StartYear:  year,
		},
	}
}

func NewEmptyHistoryRow(playerId, leagueId, leagueName, teamId, teamName string, year int) *PHistoryDto {
	return &PHistoryDto{
		PlayerId: playerId,
		Stats:    serialisePHistoryStats(NewEmptyHistorySubRow(leagueId, leagueName, teamId, teamName, year)),
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
