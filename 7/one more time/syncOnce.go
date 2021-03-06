package main

import (
	"fmt"
	"sync"
)

func main() {
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
