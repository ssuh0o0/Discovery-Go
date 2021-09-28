// interface

package main

// 고는 자바와 달리 인터페이스 내의 메서드들을 구현만 하면 그 인터페이스를 구현하는 것이 된다.
// 아까 8.3.1_test.go (다형성) 예제에서 이 코드는 다른 사람 코드에 있다고 생각하자.

type Triangle struct {
	Width, Height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Width * t.Height
}

// 그럼 이 부분을 제외한 부분만 구현해주면 된다.
// Triangle 코드를 고치지 않아도 내가 만든 인터페이스를 구현하게 되는 것.
