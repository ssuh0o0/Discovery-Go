package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(0)
				fmt.Println(n, i)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		c <- i
	}
	fmt.Println("Num : ", runtime.NumCPU())
	close(c)
}
