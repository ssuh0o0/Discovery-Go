package main

import (
	"time"
)

type status int

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

const (
	UNKNOWN status = iota
	TODO
	DONE
)

// OverDue returns true if the deadline is before the current time.
func (d *Deadline) OverDue() bool {
	return d != nil && time.Time((*d).Time).Before(time.Now())
}

// OverDue returns true if the deadline is before the current time.
func (t *Task) OverDue() bool {
	return t.Deadline.OverDue()
}

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

// JSON 태그
// type DueTask struct {
// 	Title    string `json:"title"`
// 	Status   status `json:"-"`
// 	Deadline *Deadline
// }
