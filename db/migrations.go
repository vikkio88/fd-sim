package db

import "gorm.io/gorm"

func migrate(g *gorm.DB) {
	g.AutoMigrate(
		&LeagueDto{}, &MatchDto{}, &ResultDto{},
		&TableRowDto{}, &RoundDto{}, &TeamDto{},
		&PlayerDto{}, &CoachDto{}, &GameDto{},
		&StatRowDto{}, &NewsDto{}, &EmailDto{},
		&PHistoryDto{}, &THistoryDto{}, &FDStatRowDto{},
		&RetiredPlayerDto{}, &FDHistoryDto{}, &LHistoryDto{},
		&DbEventDto{}, &TrophyDto{}, &OfferDto{},
	)
}
