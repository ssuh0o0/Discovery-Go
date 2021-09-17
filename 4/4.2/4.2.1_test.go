package main

import "fmt"

func Example_funcLiteral() {
	func() {
		fmt.Println("Hello")
	}()
	//Output:
	//Hello
}

func Example_funcLiteralVar() {
	printHello := func() {
		fmt.Println("Hello")
	}
	printHello()
	//Output:
	//Hello
}
