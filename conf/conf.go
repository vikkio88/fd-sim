package conf

import "time"

const (
	BigInt = 1_000_000_000

	PlayerAgeRange = "16..40"
	CoachAgeRange  = "30..80"
	SimDaySpeedMs  = 100

	DateFormatGame  = "Mon, 02 January 2006"
	DateFormatShort = "02/01/2006"

	StartingFame = 65

	StatsRowsLimit = 10
	// Placeholder for the link in the body of a notification
	LinkBodyPH = "LINK"

	//TODO: change this to 1st July
	StartingDateMonth = time.August
	StartingDateDay   = 20
	//
	SummerMarketWindowStart = time.July
	SummerMarketWindowEnd   = time.August
	WinterMarketWindowStart = time.January
	WinterMarketWindowEnd   = time.January
)
