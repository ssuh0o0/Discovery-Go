package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	s := even(sum, nums...)
	fmt.Println(s)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, s ...int) int {
	var ss []int
	for _, v := range s {
		if v%2 == 0 {
			ss = append(ss, v)
		}
	}

	evenSum := f(ss...)
	return evenSum
}
