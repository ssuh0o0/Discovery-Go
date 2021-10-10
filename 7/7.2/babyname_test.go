package main

import (
	"fmt"
)

func BabyNameGenerator(first string, second string) func() string {
	next, fi, se := "", 0, 0
	return func() string {
		if se == len(second) && fi == len(first) {
			next = ""
			return next
		}
		if se == len(second) {
			fi++
			se = 0
		}
		se++
		next = string(first[fi]) + string(second[se])
		return next
	}
}

func ExampleBabyGen() {
	baby := BabyNameGenerator("준경수희", "지형현기")
	for _, val := range baby() {
		fmt.Println(string(val))
	}
	//output: .
}
