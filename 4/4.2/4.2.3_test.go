package main

import (
	"fmt"
	"strings"
)

// 클로저는 외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드
func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	//output:
	//[bill tom jane]
}
