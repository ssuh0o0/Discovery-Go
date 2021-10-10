// 7.2.1 일대일 단방향 채널 소통
package main

import "fmt"

// func Example_channel() {
// 	// 새 정수 채널 만들기 c1
// 	c1 := make(chan int)
// 	// c1을 c2 에 복사 , 둘은 같은 채널이 됨
// 	var c2 chan int = c1
// 	// c3는 받기 전용 채널
// 	var c3 <-chan int = c1
// 	// c4는 보내기 전용 채널
// 	var c4 chan<- int = c1
// 	fmt.Println(c2)
// 	fmt.Println(c3)
// 	fmt.Println(c4)

// }

func Example_simpleChannel() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 1
		c <- 2
		c <- 3
	}()
	for num := range c {
		fmt.Println(num)
	}
	//output:
	//1
	//2
	//3
}

func Example_simpleChannel2() {
	ch := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for num := range ch {
		fmt.Println(num)
	}
	//output:
	//1
	//2
	//3
}
