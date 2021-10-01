package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaeyeom/gogo/task"
)

var m = task.NewInMemoryAccessor()

type ResponseError struct {
	Err error
}

func check() {
	if err != nil {
		log.Println(err)
		return
	}
}

type Response struct {
	ID    task.ID       `json:"id,omitempty"`
	Task  task.Task     `json:"task"`
	Error ResponseError `json:"error"`
}

func getTasks(r *http.Request) ([]task.Task, error) {
	var result []task.Task
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	encodedTasks, ok := r.PostForm["task"]
	if !ok {
		return nil, errors.New("task parameter expected")
	}
	for _, encodedTask := range encodedTasks {
		var t task.Task
		if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func apiGetHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	t, err := m.Get(id)
	err = json.NewEncoder(w).Encode(Response{
		ID:    id,
		Task:  t,
		Error: ResponseError{err},
	})
	check()
}

func apiPutHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	tasks, err := getTasks(r)
	check()
	for _, t := range tasks {
		err = m.Put(id, t)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Task:  t,
			Error: ResponseError{err},
		})
		check()
	}
}

func apiPostHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := getTasks(r)
	check()
	for _, t := range tasks {
		id, err := m.Post(t)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Task:  t,
			Error: ResponseError{err},
		})
		check()
	}
}

func apiDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	err := m.Delete(id)
	err = json.NewEncoder(w).Encode(Response{
		ID:    id,
		Error: ResponseError{err},
	})
	check()
}
