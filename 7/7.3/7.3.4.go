// 7.3.4 분산처리

package main

import "sync"

type IntPipe func(<-chan int) <-chan int

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

func Distribute(p IntPipe, n int) IntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}

// Distribute와 Chain을 이용해보자

out := Chain(Cut, Distribute(Chain(Draw, Paint, Decorate), 10), Box) (in)
// in 으로 들어온 자료가 하나의 고루틴 cut을 거쳐 10개로 나눠진 Draw, Paint, Decorate가 일렬로 연결된 파이프라인을 거쳐
// Box 고루틴하나로 합쳐진다.

out := Chain(Cut, Distribute(Draw, 6), Distribute(Paint, 10), Distribute(Decorate, 3), Box)(in)
// 위와 같은 out은 cut 고루틴을 거친 자료들이 6개의 Draw고루틴으로 분산되고, 
// 다시 합쳐져서, 10개의 Paint 고루틴으로 분산되고,
// 다시 합쳐져서, 3개의 Decorate 고루틴으로 분산되고,
// 다시 합쳐져서, Box 고루틴을 거쳐서 out으로 나온다.

