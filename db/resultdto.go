package db

type ResultDto struct {
	MatchId     string
	Match       MatchDto `gorm:"foreignKey:MatchId"`
	GoalsHome   int
	GoalsAway   int
	ScorersHome string
	ScorersAway string
}
