package models

const (
	WinPoints  = 3
	DrawPoints = 1
)

type Row struct {
	team         TPH
	Played       int
	Wins         int
	Draws        int
	Losses       int
	Score        int
	GoalScored   int
	GoalConceded int
}

func newRow(team *Team) *Row {
	return &Row{
		team: team.PH(),
	}
}

func (r *Row) UpdateGoals(scored, conceded int) {
	r.GoalScored += scored
	r.GoalConceded += conceded
}

func (r *Row) AddWin() {
	r.Played++
	r.Wins++
	r.Score += WinPoints
}

func (r *Row) AddLoss() {
	r.Played++
	r.Losses++
}

func (r *Row) AddDraw() {
	r.Played++
	r.Draws++
	r.Score += DrawPoints
}

type Table struct {
	order []string
	rows  map[string]*Row
	count int
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

func (t *Table) updateTableOrder() {}

func (t *Table) Rows() []*Row {
	rows := make([]*Row, t.count)
	for i, id := range t.order {
		rows[i] = t.rows[id]
	}

	return rows
}
