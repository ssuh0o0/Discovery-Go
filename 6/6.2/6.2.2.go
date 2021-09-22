package main

import (
	"errors"
	"fmt"

	"github.com/jaeyeom/gogo/task"
)

type ID string

type DataAccess interface {
	Get(id ID) (task.Task, error)
	Put(id ID, t task.Task) error
	Post(t task.Task) (ID, error)
	Delete(id ID) error
}

type MemoryDataAccess struct {
	tasks  map[ID]task.Task
	nextID int64
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]task.Task{},
		nextID: int64(1),
	}
}

var ErrTaskNoExist = errors.New("task does not exist")

func (m *MemoryDataAccess) Get(id ID) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return task.Task{}, ErrTaskNoExist
	}
	return t, nil
}

func (m *MemoryDataAccess) Put(id ID, t task.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNoExist
	}
	m.tasks[id] = t
	return nil
}

func (m *MemoryDataAccess) Post(t task.Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNoExist
	}
	delete(m.tasks, id)
	return nil
}
