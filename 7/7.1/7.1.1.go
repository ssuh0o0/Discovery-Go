// 고루틴
// 7.1.1 병렬성과 병행성
// 병렬성 : 두 사람이 각각 다른 업무를 보고 있는 것, 물리적으로 별개로 업무를 수행함.
// 병행성 : 커피를 마시며, 신문을 보는 것, 동시에 두 가지를 하고 있지만, 물리적으로 두 흐름이 있지는 않음.

// go 키워드를 쓰면 고루틴이 됨 -> 현재 흐름과는 상관없는 흐름이 됨.
// 누가 먼저 수행될지, go루틴은 실행될지 안될지도 모름.

package main

import "fmt"

func main() {
	go func() {
		fmt.Println("In goroutine")
	}()
	fmt.Println("In main routine")
}
