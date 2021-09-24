// 7.4.2 atomic과 sunc.WaitGroup
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cnt := int64(10)
	for i := 0; i < 10; i++ {
		go func() {
			//do someting
			cnt-- // !!
		}()
	}
	for cnt > 0 { // !!
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println(cnt)
}

// !! 부분에서 경쟁상태가 발생한다.
// 이때, sync/atomic 패키지는 이러한 경쟁상태를 없애주는 데 사용된다.
// cnt-- 을 atomic.AddInt64(&cnt, -1) 로
// cnt>0 을 atomic.LoadInt64(&cnt) > 0 으로 바꿔주면 된다.

// 아래 코드는 채널의 싱크를 맞춰주어 경쟁상태를 없앴다.
// 그러나 읽은 사람이 이것이 정확히 무엇을 하고자 하는 것인지 바로 파악하기 어려울 수 있다.
func main2() {
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

// 따라서 WaitGroup을 써준다.

func main3() {
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
