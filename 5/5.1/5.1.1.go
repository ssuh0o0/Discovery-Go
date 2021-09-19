package main

import "time"

//1.
var task = struct {
	title string
	done  bool
	due   *time.Time
}{"laudry", false, nil}

//2.
type Task struct {
	title string
	done  bool
	due   *time.Time
}

// 새로운 변수는 이렇게 선언가능하다.

var myTask1 Task
var myTask2 = Task{"laundry", false, nil}
var myTask3 = Task{title: "laundry"}
var myTask4 = Task{title: "laundry", done: true}
var myTask5 = Task{
	title: "laundry",
	done:  true,
}
