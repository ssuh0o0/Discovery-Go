package main

import "fmt"

func Example_sort() {
	list := []int{3, 5, 1, 4, 9}
	fmt.Println(sort(list))
	//Output:
	//[1 3 4 5 9]
}
