package main

import "fmt"

func main() {
	N := 0
	fmt.Scanln(&N)
	for i := 1; i < N+1; i++ {
		fmt.Printf("타잔이 %d원 짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다.\n", (i)*10, (i+1)*10)
	}
}
