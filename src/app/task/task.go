package task

import "time"

type Task struct {
	ID               string
	Name             string
	RunType          int // 1=Immediate, 2=schedule
	RunOn            time.Time
	TaskDependencyID []string
	Status           bool
	Execute          func() int
	Result           []interface{}
}
