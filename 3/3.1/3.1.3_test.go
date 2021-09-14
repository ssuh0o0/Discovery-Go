package main

import "fmt"

func Example_modifyBytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	//output:
	//각나다
}
