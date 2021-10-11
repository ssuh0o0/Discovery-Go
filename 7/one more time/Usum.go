package main

import (
	"context"
	"fmt"
	"strconv"
)

func Ucheck(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 0, 1)
		for {
			select {
			case i := <-c:
				c = USum(i)(ctx, c)
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

type IntPipe func(context.Context, <-chan int) <-chan int

func USum(n int) IntPipe {
	aMap := make(map[int]int)
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)

			for val := range in {
				sum := 0
				num := strconv.Itoa(val)
				nArr := []rune(num)
				for _, i := range nArr {
					a, _ := strconv.Atoi(string(i))
					sum += int(a)
				}

				_, ok := aMap[sum]
				if ok {
					continue
				} else {
					aMap[sum] = val
				}

				select {
				case out <- aMap[sum]:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for p := range Ucheck(ctx) {
		if p > 1000 {
			break
		}
		fmt.Println(p)
	}
}
