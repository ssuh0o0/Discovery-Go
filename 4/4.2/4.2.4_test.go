package main

import "fmt"

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	//output:
	//1 2 3 4 5
	//6 7 8 9 10
}

func ExampleNewIntGenerator_multiple() {
	gen1 := NewIntGenerator()
	gen2 := NewIntGenerator()
	fmt.Println(gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen1(), gen1())
	//output:
	//1 2 3
	//1 2 3 4 5
	//4 5 6
}
