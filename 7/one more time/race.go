package main

import (
	"fmt"
	"sync"
)

func main1() {
	req, resp := make(chan struct{}), make(chan int64)
	cnt := int64(10)
	go func(cnt int64) {
		defer close(resp)
		for _ = range req {
			cnt--
			resp <- cnt
		}
	}(cnt)
	for i := 0; i < 10; i++ {
		go func() {
			//do something
			req <- struct{}{}
		}()
	}
	for cnt = <-resp; cnt > 0; cnt = <-resp {

	}
	close(req)
	fmt.Println(cnt)
}

func main2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//Do something
		}()
	}
	wg.Wait()
}
