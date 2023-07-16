package db

import (
	"fdsim/models"
	"fdsim/utils"

	"gorm.io/gorm"
)

type TeamRepo struct {
	g *gorm.DB
}

func (repo *TeamRepo) TableRow(id string) *models.TPHRow {
	var dto TableRowIndexDto
	repo.g.Raw(`SELECT team_id, played, wins, draws, losses, points, goal_scored, goal_conceded, position
	FROM (
		SELECT team_id, played, wins, draws, losses, points, goal_scored, goal_conceded,
			   ROW_NUMBER() OVER (ORDER BY points DESC, goal_scored DESC, goal_conceded ASC) AS position
		FROM table_row_dtos
	) AS subquery
	WHERE team_id = ?`, id).Find(&dto)

	return dto.TPHRow()
}

func NewTeamsRepo(g *gorm.DB) *TeamRepo {
	return &TeamRepo{
		g,
	}
}

func (tr *TeamRepo) InsertOne(t *models.Team) {
	tdto := DtoFromTeam(t)

	tr.g.Create(&tdto)
}

func (tr *TeamRepo) Insert(teams []*models.Team) {
	tdtos := make([]TeamDto, len(teams))
	for i, t := range teams {
		tdtos[i] = DtoFromTeam(t)
	}

	tr.g.Create(&tdtos)
}

func (tr *TeamRepo) ById(id string) (*models.TeamDetailed, bool) {
	var t TeamDto
	trx := tr.g.Model(&TeamDto{}).Preload(playersRel).
		// Tried to db order players but it wont work
		/*, func(db *gorm.DB) *gorm.DB {
			return db.Order(`skill DESC`)
		}*/
		Preload("Coach").
		Preload("History").
		Find(&t, "Id = ?", id)

	if trx.RowsAffected != 1 {
		return nil, false
	}

	return t.TeamDetailed(), true
}

func (tr *TeamRepo) OneByFame(fame utils.Perc) *models.TPH {
	var t TeamDto
	tr.g.Raw(`select t.id, t.name
	from team_dtos t
	LEFT JOIN player_dtos p
	WHERE p.team_id  = t.id GROUP BY (p.team_id) HAVING avg(p.skill) <= ?
	ORDER by RANDOM()  limit 1`, fame.Val()).Scan(&t)

	return t.TeamPH()
}

func (tr *TeamRepo) GetRandom() *models.TPH {
	var t TeamDto
	tr.g.Order("RANDOM()").Limit(1).First(&t)
	return t.TeamPH()
}

func (tr *TeamRepo) Truncate() {
	tr.g.Where("1 = 1").Delete(&TeamDto{})
}

func (tr *TeamRepo) DeleteOne(id string) {
	tr.g.Where("Id = ?", id).Delete(&TeamDto{})
}

func (tr *TeamRepo) Delete(ids []string) {
	tr.g.Delete(&TeamDto{}, ids)
}

func (tr *TeamRepo) Count() int64 {
	var c int64
	tr.g.Model(&TeamDto{}).Count(&c)

	return c
}

func (tr *TeamRepo) All() []*models.Team {
	var tdtos []TeamDto
	tr.g.Model(&TeamDto{}).Preload(playersRel).Find(&tdtos)

	ts := make([]*models.Team, len(tdtos))
	for i, t := range tdtos {
		ts[i] = t.Team()
	}
	return ts
}
