package main

import "fmt"

func Example_simpleChannel1() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	//output:
	//1
	//2
}

func Example_simpleChannel2() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		defer close(c)
	}()
	for val := range c {
		fmt.Println(val)
	}
	//output:
	//1
	//2
}

func Example_simpleChannel3() {
	ch := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
		}()
		return c
	}()

	for val := range ch {
		fmt.Println(val)
	}
	//output:
	//1
	//2
}
