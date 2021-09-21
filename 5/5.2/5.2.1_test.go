package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Deadline struct {
	time.Time
}

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// JSON 직렬화
// JSON패키지는 대문자로 시작하는 필드들만 JSON으로 직렬화한다.
func Example_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	// JSON 데이터를 인코딩하기 위해 Marshal함수를 사용합니다
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	//output:
	// {"Title":"Laundry","Status":2,"Deadline":1439739780}
}

// {"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}

// JSON 역직렬화
// 여기서는 JSON이 주어지고, Task구조체의 값들을 채워넣을 수 있어야함.
func Example_unmarshalJSON() {
	//b := []byte(`{"Title":"Buy Milk","Status":2,"Deadline":"2015-08-16T15:43:00Z"}`)
	b := []byte(`{"Title":"Buy Milk","Status":2,"Deadline":1439739780}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	//output:
	//Buy Milk
	//2
	//2015-08-16 15:43:00 +0000 UTC
}

//JSON 태그
//JSON상에서 태그 이름을 바꾸고 싶거나, 출력하길 원치 않는 경우
type MyStruct struct {
	Title string `json:"title"`
	// 출력을 원치 않음
	Internal status `json:"-"`
	// 결과가 0인 경우에는 출력을 하지 않음
	Value int64 `json:",omitempty"`
	//json에서는 문자열로 나타남
	ID int64 `json:",string"`
}

//Json 직렬화 사용자 정의
// func (s status) MarshalJSON() ([]byte, error) {
// 	switch s {
// 	case UNKNOWN:
// 		return []byte(`"UNKNOWN`), nil
// 	case TODO:
// 		return []byte(`"TODO`), nil
// 	case DONE:
// 		return []byte(`"DONE`), nil
// 	default:
// 		return nil, errors.New("status.MarshalJSON: unknown value")
// 	}
// }

// func (s *status) UnmarshalJSON(data []byte) error {
// 	switch string(data) {
// 	case `"UNKNOWN"`:
// 		*s = UNKNOWN
// 	case `"TODO"`:
// 		*s = TODO
// 	case `"DONE"`:
// 		*s = DONE
// 	default:
// 		return errors.New("status.UnmarshalJSON: unknown value")
// 	}
// 	return nil
// }

//Unix() 사용

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

//구조체가 아닌 자료형 처리

//Json에서 사용하는 map은 모두 문자열 이어야함.
// func Example_mapMarshalJSON() {
// 	b, _ := json.Marshal(map[string]string{
// 		"Name": "John",
// 		"Age":  "16",
// 	})
// 	fmt.Println(string(b))
// 	//output:
// 	//{"Age":"16","Name":"John"}
// }

// 아무 자료형이나 쓰려면 interface{} 쓰면됨.
func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]interface{}{
		"Name": "John",
		"Age":  16,
	})
	fmt.Println(string(b))
	//output:
	//{"Age":16,"Name":"John"}
}

// JSON 필드 조작하기
type Fields struct {
	VisibleField   string `json:"visibleField"`
	InvisibleField string `json:"invisibleField"`
}

func ExampleOmitField() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:"invisibleField,omitempty"`
		Additional     string `json:"additional,omitempty"`
	}{Fields: f, Additional: "c"})
	fmt.Println(string(b))
	//output: {"visibleField":"a","additional":"c"}
}
