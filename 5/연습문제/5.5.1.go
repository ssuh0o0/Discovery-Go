package main

import (
	"fmt"
	"time"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Deadline struct {
	time.Time
}

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix+"  ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}

func (t *Task) MarkDone() {
	t.Status = DONE
	for i := range t.SubTasks {
		t.SubTasks[i].MarkDone()
	}
}
