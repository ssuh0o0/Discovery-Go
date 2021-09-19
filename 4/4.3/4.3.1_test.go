package main

import "fmt"

type VertexID int

func (id VertexID) String() string {
	return fmt.Sprintf("VertexID(%d)", id)
}

func ExampleVertexID_String() {
	i := VertexID(100)
	fmt.Println(i)
	//output:
	//VertexID(100)
}
