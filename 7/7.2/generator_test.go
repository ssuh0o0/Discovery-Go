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

func FibonacciGenerator(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return
		}
		return next
	}
}

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, ",")
	}
	//output: .
}

func ExampleFibonacciGenerator() {
	g := FibonacciGenerator(15)

	for v := g(); a > 0; v = g() {
		fmt.Println(v)
	}
}
