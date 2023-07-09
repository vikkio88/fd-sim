package db

import (
	"encoding/json"
	"fdsim/models"
	"fmt"
	"time"
)

type THistoryDto struct {
	TeamId string `gorm:"primarykey"`
	Stats  string

	Team TeamDto
}

func (p *THistoryDto) Update(row TableRowIndexDto, leagueId, leagueName string, gameDate time.Time) {
	existingStats := unserialiseTHistoryStats(p.Stats)
	hstat := models.NewTHistoryRow(row.TPHRow(), leagueId, leagueName, gameDate)
	existingStats = append(existingStats, newSubRowFromTHistory(hstat))

	p.Stats = serialiseTHistoryStats(existingStats)
}

type THistorySubRow struct {
	LeagueId      string
	LeagueName    string
	Played        int
	Wins          int
	Draws         int
	Losses        int
	Points        int
	GoalScored    int
	GoalConceded  int
	Year          int
	FinalPosition int
}

func newSubRowFromTHistory(row *models.THistoryRow) THistorySubRow {
	return THistorySubRow{
		LeagueId:      row.LeagueId,
		LeagueName:    row.LeagueName,
		Played:        row.Played,
		Wins:          row.Wins,
		Draws:         row.Draws,
		Losses:        row.Losses,
		Points:        row.Played,
		GoalScored:    row.Played,
		GoalConceded:  row.Played,
		FinalPosition: row.FinalPosition,
		Year:          row.Year,
	}
}

func newSubRowsFromTHistory(h *models.THistoryRow) []THistorySubRow {
	return []THistorySubRow{
		newSubRowFromTHistory(h),
	}
}

func DtoFromTHistoryRow(h *models.THistoryRow) THistoryDto {
	return THistoryDto{
		TeamId: h.TeamId,
		Stats: serialiseTHistoryStats(
			newSubRowsFromTHistory(h),
		),
	}
}

func serialiseTHistoryStats(tHistorySubRow []THistorySubRow) string {
	var result string
	data, _ := json.Marshal(tHistorySubRow)
	result = string(data)

	return result
}

func unserialiseTHistoryStats(s string) []THistorySubRow {
	if s == "" {
		return []THistorySubRow{}
	}

	var result []THistorySubRow
	data := s
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("error while loading", err)
		return []THistorySubRow{}
	}

	return result
}