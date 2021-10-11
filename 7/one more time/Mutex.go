package main

import "sync"

type Resource []interface{}

type Accessor struct {
	R *Resource
	L *sync.Mutex
}

func (acc *Accessor) Use() {
	acc.L.Lock()
	acc.L.Unlock()
}

type ConcurrentMap struct {
	M map[string]string
	L *sync.RWMutex
}

func (m ConcurrentMap) Get(key string) string {
	m.L.RLock()
	defer m.L.RUnlock()
	return m.M[key]
}

func (m ConcurrentMap) Set(key, value string) {
	m.L.Lock()
	m.M[key] = value
	m.L.Unlock()
}

func main() {
	m := ConcurrentMap{map[string]string{}, &sync.RWMutex{}}
}
