// 7.3.9 동적으로 고루틴 이어붙이기
package main

import "context"

type IntPipe func(context.Context, <-chan int) <-chan int

func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// n의 배수를 걸러내는 파이프라인을 반환한다.
func FilterMultiple(n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

// Primes 무한소수 생성기
func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)
		for {
			select {
			case i := <-c:
				c = FilterMultiple(i)(ctx, c)
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}
