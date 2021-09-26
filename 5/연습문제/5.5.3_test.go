package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type status int

type Task struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}

func Example_unmarshalJSON() {
	b := []byte(`{"ID": "6" , "Name" : "soo"}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.ID)
	fmt.Println(t.Name)
	//output:
	//6
	//soo
}
