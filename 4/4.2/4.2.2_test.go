package main

import (
	"fmt"
	"strings"
)

func ExampleReadFrom_Print() {
	r := strings.NewReader("bill\ntom\njane\n")
	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
	//Output:
	//( bill )
	//( tom )
	//( jane )

}
