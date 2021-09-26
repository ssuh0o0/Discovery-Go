package main

import (
	"fmt"
)

func ExampleIncludeSubTasks_MarkDone() {
	markTask := Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}
	markTask.MarkDone()
	fmt.Println(IncludeSubTasks(markTask))
	// Output:
	// [v] Laundry <nil>
	//   [v] Wash <nil>
	//     [v] Put <nil>
	//     [v] Detergent <nil>
	//   [v] Dry <nil>
	//   [v] Fold <nil>
}
