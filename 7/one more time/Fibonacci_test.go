package main

import "fmt"

func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func FibGenerator(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}

func Example_Fibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Println(fib)
	}
	//output:
	// .
}

func Example_FibGenerator() {
	fi := FibGenerator(15)
	for fib := fi(); fib >= 0; fib = fi() {
		fmt.Println(fib)
	}
	//output:
	//.
}
