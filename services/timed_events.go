package services

import (
	"fdsim/conf"
	"fdsim/utils"
	"time"
)

type marketCheck struct {
	openingDate bool
	closingDate bool
	summer      bool
	winter      bool
	opening     string
	closing     string
}

func (m marketCheck) isOpen() bool {
	return m.openingDate && (m.summer || m.winter)
}

func makeMarketWindows(date time.Time) []time.Time {
	thisYear := date.Year()

	return []time.Time{
		utils.NewDate(thisYear, conf.SummerMarketWindowStart, 1),
		utils.NewDate(thisYear, conf.SummerMarketWindowEnd, 31),
		utils.NewDate(thisYear, conf.WinterMarketWindowStart, 1),
		utils.NewDate(thisYear, conf.WinterMarketWindowEnd, 31),
	}
}

func calculateTransferWindowDates(date time.Time) marketCheck {
	dates := makeMarketWindows(date)

	if date.Equal(dates[0]) {
		return marketCheck{
			openingDate: true, summer: true,
			opening: dates[0].Format(conf.DateFormatShort),
			closing: dates[1].Format(conf.DateFormatShort),
		}
	}

	if date.Equal(dates[1]) {
		return marketCheck{closingDate: true, summer: true,
			opening: dates[0].Format(conf.DateFormatShort),
			closing: dates[1].Format(conf.DateFormatShort),
		}
	}

	if date.Equal(dates[2]) {
		return marketCheck{openingDate: true, winter: true,
			opening: dates[2].Format(conf.DateFormatShort),
			closing: dates[3].Format(conf.DateFormatShort)}
	}

	if date.Equal(dates[3]) {
		return marketCheck{closingDate: true, winter: true,
			opening: dates[2].Format(conf.DateFormatShort),
			closing: dates[3].Format(conf.DateFormatShort),
		}
	}

	if date.After(dates[0]) && date.Before(dates[1]) {
		return marketCheck{summer: true,
			opening: dates[2].Format(conf.DateFormatShort),
			closing: dates[3].Format(conf.DateFormatShort)}
	}

	if date.After(dates[2]) && date.Before(dates[3]) {
		return marketCheck{winter: true}
	}

	return marketCheck{}
}
