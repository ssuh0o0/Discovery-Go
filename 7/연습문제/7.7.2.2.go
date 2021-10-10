// 클로저

package main

import "fmt"

func PlusNames(first, second string) func() string {
	ss := "ss"
	return func() string {
		for _, f := range first {
			for _, s := range second {
				ss = string(f) + string(s)
			}
		}
		return ss
	}
}

func main() {
	bname := PlusNames("성정명재경", "준호우훈진")
	for name := bname(); name != ""; name = bname() {
		fmt.Println(name)
	}

}
