// 7.3.7 컨텍스트 활용하기

package main

import (
	"context"
	"fmt"
)

func PlusOne(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-ctx.done:
				return
			}
		}
	}()
	return out
}

func main() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	ctx, cancel := context.WithCancel(context.Background()) //ctx에는 새로 생성된 컨텍스트가, cancel에는
	nums := PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			cancel()
			break
		}
	}
}

// 밖에서 취소신호를 받을 수 있는 경우는 context.Context로 넘겨받을수 있다ㅏ.
// context.Context는 계층구조로 되어있다.
// context.Background()가 가장 상위에 있다. 따라서, 프로그램이 끝날 때까지 취소되지 않고 살아있다.
// 상위구조가 취소되면 하위에 있는 모든 컨텍스트도 취소된다.
// WithDeadline, WithTimeout을 이용해 만든 ctx를 이용해 호출하면 시간이 지나면 취소되게 만들 수 있다.
// WithValue를 이용하면 인증 토큰 같이 요청 범위 내에 있는 값들을 보낼 수 있어 편리하다.
