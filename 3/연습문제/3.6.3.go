package main

import "fmt"

func isInList(v string, list []string) bool {
	low := 0
	high := len(list) - 1
	mid := 0

	for low <= high {
		mid = (low + high) / 2

		if list[mid] == v {
			return true
		} else if list[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	list := []string{"a", "b", "x"}
	fmt.Println(isInList("b", list))
	fmt.Println(isInList("z", list))
}
