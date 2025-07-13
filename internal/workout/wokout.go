package workout

import "time"

type Workout struct {
	ID          string
	Name        string
	Date        time.Time
	description string
	TimeSpent   string
}

type Task struct {
}

type TaskType struct {
	name        string
	description string
}
