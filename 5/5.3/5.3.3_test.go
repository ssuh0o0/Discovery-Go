package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

// ***** 정렬 인터페이스의 구현 *****
type Interface interface {
	Len() int
	// i, j 비교해서 작은 것이 앞으로 가야함.
	Less(i, j int) bool
	Swap(i, j int)
}

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) || (strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	//output:
	//[AppStore iPad iPhone MacBook]
}

// ***** 정렬 알고리즘 *****
// * sort.Sort는 퀵정렬을 쓴다.
// * 그 중에서도 피벗 3개를 골라서 가운데 값을 고르는 중위법을 사용한다.
// * 너무 깊은 퀵정렬에 들어가게 되면, 힙 정렬을 사용한다.
// * 힙 정렬을 사용할때도 7개 이하의 자료에서는 삽입정렬을 사용한다.

// ***** 힙 *****
// 힙은 가장 작은 원소를 O(logN)의 시간 복잡도로 꺼낼 수 있는 자료구조.
// 그러나 퀵정렬보다 대체적으로 느리고, 캐시를 효율적으로 사용할 수 없음.
// o(nlogN)
// 작은(큰) 값 부터 소비해야 하거나,정렬된 여러자료를 하나로 합치는 경우에 효율적이다.
type heap_Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		fmt.Println(heap.Pop(&apple))
	}

	//output:
	//AppStore
	//iPad
	//iPhone
	//MacBook
}
