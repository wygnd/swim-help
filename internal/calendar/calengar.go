package calendar

import "time"

const rangeIndex = 8

func Today() time.Time {
	return time.Now()
}

func RangeDates(prev bool) [rangeIndex]time.Time {
	var dates [rangeIndex]time.Time
	today := Today()

	dates[0] = today
	for i := 1; i < rangeIndex; i++ {
		if prev {
			dates[i] = dates[i-1].AddDate(0, 0, -1)
		} else {
			dates[i] = dates[i-1].AddDate(0, 0, 1)
		}
	}

	return dates
}
