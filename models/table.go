package models

import (
	"fmt"
	"sort"
)

const (
	WinPoints  = 3
	DrawPoints = 1
)

// TeamPlaceholder and Table Row Wrapper
type TPHRow struct {
	Index int
	Team  TPH
	Row   *Row
}

type Row struct {
	Team         string
	Played       int
	Wins         int
	Draws        int
	Losses       int
	Points       int
	GoalScored   int
	GoalConceded int
}

func newRow(team *Team) *Row {
	return &Row{
		Team: team.Id,
	}
}

func (r *Row) String() string {
	return fmt.Sprintf("%s\tw: %d d: %d l: %d , gs: %d gc: %d , %d",
		r.Team, r.Wins, r.Draws, r.Losses, r.GoalScored,
		r.GoalConceded, r.Points,
	)
}

func (r *Row) UpdateGoals(scored, conceded int) {
	r.GoalScored += scored
	r.GoalConceded += conceded
}

func (r *Row) AddWin() {
	r.Played++
	r.Wins++
	r.Points += WinPoints
}

func (r *Row) AddLoss() {
	r.Played++
	r.Losses++
}

func (r *Row) AddDraw() {
	r.Played++
	r.Draws++
	r.Points += DrawPoints
}

type Table struct {
	order []string
	rows  map[string]*Row
	count int
}

func NewTableFromRows(rows []*Row) *Table {
	count := len(rows)
	order := make([]string, len(rows))
	rowsMap := map[string]*Row{}
	for i, r := range rows {
		rowsMap[r.Team] = r
		order[i] = r.Team
	}
	t := &Table{
		order: order,
		rows:  rowsMap,
		count: count,
	}

	t.updateTableOrder()

	return t
}

func NewTable(teams []*Team) *Table {
	count := len(teams)
	rows := map[string]*Row{}
	order := make([]string, len(teams))
	for i, t := range teams {
		rows[t.Id] = newRow(t)
		order[i] = t.Id
	}
	return &Table{
		order: order,
		rows:  rows,
		count: count,
	}
}

// returns team_id at index
func (t *Table) Get(index int) (string, *Row) {
	tId := t.order[index]
	return tId, t.rows[tId]
}

func (t *Table) Update(round *Round) {
	res, ok := round.Results()
	if !ok {
		return
	}

	for id, r := range res {
		m := round.MatchMap[id]
		switch r.X12() {
		case R1:
			{
				t.rows[m.Home.Id].AddWin()
				t.rows[m.Away.Id].AddLoss()
			}
		case R2:
			{
				t.rows[m.Away.Id].AddWin()
				t.rows[m.Home.Id].AddLoss()
			}
		case RX:
			{
				t.rows[m.Home.Id].AddDraw()
				t.rows[m.Away.Id].AddDraw()
			}
		}
		t.rows[m.Home.Id].UpdateGoals(r.GoalsHome, r.GoalsAway)
		t.rows[m.Away.Id].UpdateGoals(r.GoalsAway, r.GoalsHome)
	}

	t.updateTableOrder()
}

func (t *Table) updateTableOrder() {
	sort.SliceStable(t.order, func(i, j int) bool {
		rowI := t.rows[t.order[i]]
		rowJ := t.rows[t.order[j]]

		if rowI.Points != rowJ.Points {
			return rowI.Points > rowJ.Points
		}

		return (rowI.GoalScored - rowI.GoalConceded) > (rowJ.GoalScored - rowJ.GoalConceded)
	})
}

func (t *Table) Len() int {
	return t.count
}

func (t *Table) Rows() []*Row {
	rows := make([]*Row, t.count)
	for i, id := range t.order {
		rows[i] = t.rows[id]
	}

	return rows
}

func (t *Table) TPHRows(mappings TeamMap) []*TPHRow {
	rows := make([]*TPHRow, len(t.order))
	for i, id := range t.order {
		rows[i] = &TPHRow{
			Index: i,
			Team:  mappings[id].PH(),
			Row:   t.rows[id],
		}
	}

	return rows
}
