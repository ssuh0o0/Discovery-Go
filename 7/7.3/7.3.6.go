// 7.3.6 파이프라인 중단하기

package main

import (
	"fmt"
	"runtime"
	"time"
)

func PlusOne(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	done := make(chan struct{})
	nums := PlusOne(done, PlusOne(done, PlusOne(done, PlusOne(done, PlusOne(done, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	for _ = range nums {
		//consume all nums
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}
