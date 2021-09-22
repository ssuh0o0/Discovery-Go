// 빈 인터페이스와 형 단언
package main

import (
	"container/heap"
	"fmt"
)

func ExampleCaseInsensitive_heapString() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		popped := heap.Pop(&apple)
		// Pop의 반환형은 interface{}라고 되어있다. 여기서는 string으로 형단언을 해줌.
		s := popped.(string)
		fmt.Println(s)
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}

// 실제 interface만 형 단언을 쓸 수 있는 것은 아니다.

// var r io.Reader = NewReader()
// f := r.(os.File)

// 이렇게 할 경우, r이 os.File형 인 경우 작동하지만, 아닐 경우 패닉이 발생한다.
// 따라서, 이렇게 해줘야 한다. ok가 false가 되는 경우 패닉이 발생하지 않는다.

// var r io.Reader = NewReader()
// f, ok := r.(os.File)
