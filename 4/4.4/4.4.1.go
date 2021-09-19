package main

import (
	"fmt"
	"time"
)

func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}

func main() {
	fmt.Println("Ladies and gentlemen!")
	CountDown(5)
}

// 비동기적 상황에서 사용되는 테크닉이 콜백.
// 콜백은 어떤 조건이 만족될 때 호출해달라고 요청하는 것으로
// 아래는 고계 함수를 이용하여 콜백을 이용함.

// time.AfterFunc(5*time.Second, func()){
// 	//메세지를 없애는 코드
// }

// time.AfterFunc(3*time.Second, func()) {
// 	fmt.Println(" I am so excited! ")
// }

// timer := AfterFunc(...)
// ...
// timer.Stop()
