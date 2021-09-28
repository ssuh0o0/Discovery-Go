// 클로저

package main

import "fmt"

func PlusNames(first, second []string) func() []string {
	var ss []string
	return func() []string {
		for _, f := range first {
			for _, s := range second {
				ss = append(ss, string(f)+string(s))
			}
		}
		return ss
	}
}

func main() {
	first := []string{"성", "정", "명", "재", "경"}
	second := []string{"준", "호", "우", "훈", "진"}

	name := PlusNames(first, second)()
	fmt.Println(name)
}
