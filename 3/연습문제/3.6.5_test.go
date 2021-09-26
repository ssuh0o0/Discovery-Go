package main

import (
	"fmt"
	"strings"
)

func NewMultiSet() map[string]int {
	return map[string]int{}
}

func Insert(m map[string]int, val string) {
	m[val]++
}

func Erase(m map[string]int, val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func Count(m map[string]int, val string) int {
	return m[val]
}

func String(m map[string]int) string {
	s := "{ "
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + "}"
}

func ExampleMultiSet() {
	m := NewMultiSet()
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "1")
	Insert(m, "2")
	Insert(m, "5")
	Insert(m, "7")
	Erase(m, "3")
	Erase(m, "5")
	fmt.Println(Count(m, "3"))
	fmt.Println(Count(m, "1"))
	fmt.Println(Count(m, "2"))
	fmt.Println(Count(m, "5"))
	//output:
	// { }
	// 0
	// { 3 3 3 3 }
	// 4
	// 3
	// 1
	// 1
	// 0
}
