package workout

import "time"

type Workout struct {
	ID        string
	Name      string
	Date      time.Time
	Desc      string
	TimeSpent string
}
