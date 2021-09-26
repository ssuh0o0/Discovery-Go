package main

import "fmt"

type Queue []interface{}

func (q *Queue) push(data interface{}) {
	*q = append(*q, data)
	fmt.Printf("push: %v\n", data)
}

func (q *Queue) pop() interface{} {
	if len(*q) == 0 {
		fmt.Println("queue is empty")
		return nil
	}
	data := (*q)[0]
	// 이렇게 짜면, q가 비었을때, 에러가 날 수 있다. 따라서, 위 처럼 조건문을 넣어주면 된다.
	*q = (*q)[1:]
	fmt.Printf("pop: %v\n", data)
	return data
}

func main() {
	q := Queue{}
	q.push(1)
	q.push(2)
	fmt.Printf("%v\n", q)

	q.pop()
	fmt.Printf("%v\n", q)
}
