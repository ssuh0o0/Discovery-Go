// 콜백
package main

import "fmt"

func PlusNames(first, second []string) []string {
	var ss []string
	for _, f := range first {
		for _, s := range second {
			ss = append(ss, string(f)+string(s))
		}
	}
	return ss
}

func BabyNames(funcP func(f, s []string) []string, n ...string) []string {
	length := len(n)
	first := n[:(length / 2)]
	second := n[(length / 2):]

	nameList := funcP(first, second)
	return nameList
}

func main() {
	names := []string{"성", "정", "명", "재", "경", "준", "호", "우", "훈", "진"}
	s := BabyNames(PlusNames, names...)
	fmt.Println(s)
}
