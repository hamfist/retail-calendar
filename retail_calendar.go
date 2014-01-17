package retail_calendar

import (
	"time"
)

func MonthRange(date time.Time) (startTime time.Time, endTime time.Time) {
	fiscalYear := date.Year()
	if date.Month() == 1 {
		fiscalYear = fiscalYear - 1
	}
	endOfJanuary := time.Date(fiscalYear, 1, 31, 0, 0, 0, 0, time.UTC)
	daysUntilSunday := 7 - endOfJanuary.Weekday()
	startTime = endOfJanuary.AddDate(0, 0, int(daysUntilSunday))
	endTime = startTime

	month := 1
	for month < 13 {
		startTime = endTime
		endTime = startTime.AddDate(0, 0, calcWeeksInMonth(month)*7)
		if date.Before(endTime) { break }
		month++
	}
	return startTime, endTime
}

// calculate number of weeks in 454 month
func calcWeeksInMonth(month int) int {
	if month == 2 || month == 5 || month == 8 || month == 11 {
		return 5
	} else {
		return 4
	}
}

