package main

import "fmt"

func BabyNames(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

func BNGenerator(first, second string) func() string {
	ss := ""
	fi, se := []rune(first), []rune(second)
	return func() string {
		for f := range fi {
			for s := range se {
				// 왜 여기 안에 return 안되지..?
				ss = string(fi[f]) + string(se[s])
			}
		}
		return ss
	}

}

func ExampleBabyNames() {
	for val := range BabyNames("가나다라마", "숑뿅뚕찡땅") {
		fmt.Print(val, ",")
	}
	//output:.
}

func Example_BNGenerator() {
	baby := BNGenerator("가나다라마", "숑뿅뚕찡땅")

	for b := baby(); b != ""; b = baby() {
		fmt.Println(b)
	}
	//output:.
}
