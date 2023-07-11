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

func (t *THistoryDto) HistoryRows() []*models.THistoryRow {
	if t.Stats == "" {
		return []*models.THistoryRow{}
	}

	stats := unserialiseTHistoryStats(t.Stats)
	if len(stats) < 1 {
		return []*models.THistoryRow{}
	}

	result := make([]*models.THistoryRow, len(stats))
	for i, s := range stats {
		result[i] = s.THistoryRow()
	}

	return result
}

func (t *THistoryDto) Update(row TableRowIndexDto, leagueId, leagueName string, gameDate time.Time) {
	existingStats := unserialiseTHistoryStats(t.Stats)
	hstat := models.NewTHistoryRow(row.TPHRow(), leagueId, leagueName, gameDate)
	existingStats = append(existingStats, newSubRowFromTHistory(hstat))

	t.Stats = serialiseTHistoryStats(existingStats)
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

func (h *THistorySubRow) THistoryRow() *models.THistoryRow {
	return &models.THistoryRow{
		LeagueId:      h.LeagueId,
		LeagueName:    h.LeagueName,
		Played:        h.Played,
		Wins:          h.Wins,
		Draws:         h.Draws,
		Losses:        h.Losses,
		Points:        h.Played,
		GoalScored:    h.Played,
		GoalConceded:  h.Played,
		FinalPosition: h.FinalPosition,
		Year:          h.Year,
	}
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
