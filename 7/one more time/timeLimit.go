package main

import (
	"fmt"
	"time"
)

func main() {
	// timelimit := time.After(5* time.Second)
	recv := make(chan int)
	send := make(chan int)
	select {
	case n := <-recv:
		fmt.Println(n)
	case send <- 1:
		fmt.Println("send 1")
	case <-time.After(5 * time.Second):
		fmt.Println("5seconds")
		return
	}
}
