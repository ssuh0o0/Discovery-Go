package main

import "fmt"

func main() {
	N := 0
	fmt.Scanln(&N)
	a, b := 0, 1
	for i := 0; i < N-2; i++ {
		a, b = b, a+b
	}
	fmt.Println(b)
}
