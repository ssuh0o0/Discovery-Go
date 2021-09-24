//7.4.3 sync.Once

package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan struct{})
	go func() { //초기화 함수
		defer close(done)
		fmt.Println("Initialized")
	}()
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) { // 고루틴이지만 <-done을 통해 받아지는 것을 기다린 후 화면에 출력
			defer wg.Done()
			<-done
			fmt.Println("Go Routine", i)
		}(i)
	}
	wg.Wait()
}

// 이렇게 어떤 코드를 한번만 수행하고자 할때 쓸 수 있는 것이 sunc.Once이다.
// 주로 분산처리를 할 때, 초기화 코드에 쓰인다.

func mainOnce() {
	var once sync.Once
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("Initialized")
			})
			fmt.Println("Go Routine", i)
		}(i)
	}
	wg.Wait()
}
