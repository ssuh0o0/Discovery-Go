// 7.2.2 생성기 패턴
package main

import "fmt"

func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, ",")
	}
	//output: 0,1,1,2,3,5,8,13,
}

// 클로저를 이용하면 이렇게 할 수 있다.

func FibonacciGenerator(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}

func ExampleFibonacciGenerator() {
	fib := FibonacciGenerator(15)
	for n := fib(); n >= 0; n = fib() {
		fmt.Print(n, ",")
	}
	//output: 0,1,1,2,3,5,8,13,
}

// 채널을 이용하면 받는 쪽에서 for의 range를 사용할 수 있다.
func BabyNames(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

func ExampleBabyNames() {
	for n := range BabyNames("성정명재경", "준호우훈진") {
		fmt.Println(n, ",")
	}
	//output:
	// 성준 ,
	// 성호 ,
	// 성우 ,
	// 성훈 ,
	// 성진 ,
	// 정준 ,
	// 정호 ,
	// 정우 ,
	// 정훈 ,
	// 정진 ,
	// 명준 ,
	// 명호 ,
	// 명우 ,
	// 명훈 ,
	// 명진 ,
	// 재준 ,
	// 재호 ,
	// 재우 ,
	// 재훈 ,
	// 재진 ,
	// 경준 ,
	// 경호 ,
	// 경우 ,
	// 경훈 ,
	// 경진 ,
}
