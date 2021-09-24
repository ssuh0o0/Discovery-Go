// 7.3.3 팬인하기

package main

import "sync"

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 실행은 이런식으로 쓰면된다.
// c:= FanIn(c1,c2,c3)
// c1, c2, c3 채널에서 나온 자료들은 모두 c로 오게된다.
