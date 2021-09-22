package main

import (
	"fmt"
)

func ExampleJoin() {
	t := Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: nil,
	}
	fmt.Println(Join(",", 1, "two", 3, t))
	//output: 1,two,3,[v] Laundry <nil>
}
