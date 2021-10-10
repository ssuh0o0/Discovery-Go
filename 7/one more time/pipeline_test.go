package main

import "fmt"

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 1
		c <- 2
	}()
	PlusTwo := Chain(PlusOne, PlusOne)(c)
	for num := range PlusTwo {
		fmt.Println(num)
	}
	//output: .
}
