package main

import (
	"fmt"
	"runtime"
	"time"
)

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

func main() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 51; i += 10 {
			c <- i
		}
	}()
	nums := PlusOne(PlusOne(PlusOne(c)))
	for num := range nums {
		fmt.Println(num)
		if num > 30 {
			break
		}
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	for _ = range nums {

	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
}
