// 7.3.1 동시성 패턴 ( 파이프라인 패턴 )
// 한 단계의 출력이 다음 단계의 입력으로 이어지는 구조
// 분업과 비슷하다고 생각하면 됨.

// 받기 전용 채널을 넘겨 받아서 입력으로 사용한다.

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range PlusOne(PlusOne(c)) {
		fmt.Printlnnum
	}

	//output:
	//7
	//5
	//10
}

// 이렇게 함수 형태에 이름을 붙이면 쉽다.
type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <- chan int {
		c := in
		for _, p in range ps {
			c = p(c)
		}
		return c
	}
}

// 위와 같은 chain을 만들면 B(A(c))를 Chain(A,B)(c)로 표기할 수 있다.