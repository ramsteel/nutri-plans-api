package date

import "time"

func GetTodayRange() (time.Time, time.Time) {
	today := time.Now()
	start := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	end := time.Date(today.Year(), today.Month(), today.Day(), 23, 59, 59, 0, today.Location())
	return start, end
}
