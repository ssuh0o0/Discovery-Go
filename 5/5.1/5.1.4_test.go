package main

import (
	"fmt"
	"time"
)

func Example_taskTestAll() {
	// d1 := Deadline(time.Now().Add(-4 * time.Hour))
	// d2 := Deadline(time.Now().Add(4 * time.Hour))
	d1 := NewDeadline(time.Now().Add(-4 * time.Hour))
	d2 := NewDeadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, d1}
	t2 := Task{"4h later", TODO, d2}
	t3 := Task{"no due", TODO, nil}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
	// Output:
	// true
	// false
	// false
}

type Address struct {
	City  string
	State string
}

type Telephone struct {
	Mobile string
	Direct string
}

type Contact struct {
	Address
	Telephone
}

func ExampleContact() {
	var c Contact
	c.Mobile = "123-456-789"
	fmt.Println(c.Telephone.Mobile)
	c.Address.City = "San Francisco"
	c.State = "CA"
	c.Direct = "N/A"
	fmt.Println(c)
	//output:
	//123-456-789
	//{{San Francisco CA} {123-456-789 N/A}}
}
