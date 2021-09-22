package main

import (
	"fmt"
	"strconv"
	"strings"
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
	Title    string
	Status   status
	Deadline *Deadline
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	t := make([]string, len(a))
	for i := range a {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String()
		}
		// if x, ok := a[i].(string); ok {
		// 	t[i] = x
		// } else if x, ok := a[i].(int); ok {
		// 	t[i] = strconv.Itoa(x)
		// } else if x, ok := a[i].(fmt.Stringer); ok {
		// 	t[i] = x.String()
		// }
	}
	return strings.Join(t, sep)
}
