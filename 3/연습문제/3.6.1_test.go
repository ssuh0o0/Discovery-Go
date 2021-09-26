package main

import "fmt"

func Example_array() {
	fruits := [4]string{"사과", "바나나", "토마토", "수박"}
	for _, fruit := range fruits {
		if HasConsonantSuffix(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)
		}
	}
	//Output:
	//사과는 맛있다.
	//바나나는 맛있다.
	//토마토는 맛있다.
	//수박은 맛있다.
}
