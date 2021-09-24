// 7.3.5 select
// select를 쓰면 여러 채널과 통신할 수 있다. ( switch와 비슷함 )

// **** 팬인하기 ****

// 아래 코드는 c1, c2, c3중 어느 채널이라도 자료가 준비되어 있으면, 그것을 바로 c로 보내는 코드이다.
select {
case n := <-c1:
	c <- n
case n := <-c2:
	c <- n
case n := <-c3:
	c <- n
}
// 만일 c1, c2, c3 중 어떤 채널이 닫혀있다면,
// 닫혀있는 채널은 막히지 않고 기본값을 계속 받아갈 수 있기 떄문에,
// 닫힌 채널이 선택되어 기본값이 받아질 가능성이 있다.
// 이거 까지 처리한 FanIn 을 만들어보자.

package main

import "fmt"

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3
	closeChan := func(c *<-chan int) bool {
		*c = nil
		openCnt--
		return openCnt == 0
	}
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			}
		}
	}()
	return out
}

func main() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	SendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}
	go SendInts(c1, 11, 14)
	go SendInts(c2, 21, 23)
	go SendInts(c3, 31, 35)
	for n := range FanIn3(c1, c2, c3) {
		fmt.Print(n, ",")
	}
}

// **** 채널 기다리지 않고 받기 ****
// 지금까지는 값이 없으면, 채널 기다리는 방식이었지만
// select를 사용하면, 채널값이 있으면 받고, 없으면 스킵하는 식으로 바꿀 수 있다.
select {
case n := <-c :
	fmt.Println(n)
default:
	fmt.Println("Data is not ready. Skipping...")
} 

// **** 시간 제한 ****
// 채널과 통신을 일정시간 동안만 기다리겠다는 time.After 함수.
select {
case n := <-recv:
	fmt.Println(n)
case send <- 1:
	fmt.Prinntln("sent 1")
case <-time.After(5 * time.Second)
	fmt.Println("No send and receive communication for 5 seconds")
	return
}

// 전체 제한시간을 걸고 싶으면 타이머 채널을 보관해두고 사용하면 된다.
timeout := time.After( 5 * time.Second)
for {
	select {
	case n := <-recv:
		fmt.Println(n)
	case send <- 1:
		fmt.Prinntln("sent 1")
	case <-timeout
		fmt.Println("5초안에 통신 안끝남")
		return
	}
}
