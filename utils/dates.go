package utils

import "time"

func GetFirstSunday(year int, month time.Month) time.Time {
	firstDay := NewDate(year, month, 1)
	dayOfWeek := firstDay.Weekday()
	daysToAdd := (7 - int(dayOfWeek) + int(time.Sunday)) % 7
	firstSunday := firstDay.AddDate(0, 0, daysToAdd)

	return firstSunday
}

func GetLastSunday(year int, month time.Month) time.Time {
	nextMonth := month + 1
	firstDayNextMonth := NewDate(year, nextMonth, 1)

	lastDay := firstDayNextMonth.AddDate(0, 0, -1)
	offset := int(lastDay.Weekday() - time.Sunday)
	if offset > 0 {
		offset = -offset
	}
	lastSunday := lastDay.AddDate(0, 0, offset)

	return lastSunday
}

func NewDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
