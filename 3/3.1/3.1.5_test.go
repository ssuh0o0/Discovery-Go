package main

import "fmt"

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	//output:
	//abcdef
	//abcdef
}

// s += "def" 대신에
// s = fmt.Sprint(s, "def")
// s = fmt.Sprint("%sdef", s)
// s = strings.Join([]string{s, "def"},"")
