package main

// 직렬화
func Example_marshalJSON() {
	t := Task{
		"Laundry",
		Done,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC()))
	}
}
b, err := json.Marshal(t)
if err != nil {
	log.Println(err)
	return
}
fmt.Println(string(b))
//output:
// {"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}

// 역직렬화
func

//JSON 태그
type DueTask struct {
	Title    string `json:"title"`
	Status   status `json:"-"`
	Deadline *Deadline
}

//Json 직렬화 사용자 정의

func (s tatus) MarshalJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN`), nil
	case TODO:
		return []byte(`"TODO`), nil
	case DONE:
		return []byte(`"DONE`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}

func (s tatus) UnmarshalJSON() (data []byte) error {
	switch string(data) {
	case UNKNOWN:
		*s = UNKNOWN
	case TODO:
		*s = TODO
	case DONE:
		*s = DONE
	default:
		return nil, errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}