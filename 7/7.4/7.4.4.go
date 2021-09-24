// 7.4.4 Mutex와 RWMutex
// 뮤텍스는 상호 배타 잠금 기능이 있다.
// 둘 이상의 고루틴에서 동시에 코드의 흐름을 제어할 수 있다.

package main

import "sync"

type Accessor struct {
	R *Resource   // 접근하고자 하는 자원 포인터
	L *sync.Mutex // 뮤텍스 포인터
}

// 생성할 떄, Acessor{&resource, &sync.Mutex{}}와 같이 할당해준다.

func (acc *Accessor) Use() {
	acc.L.Lock()
	acc.L.Unlock()
}

// sunc.RWMutex는 좀 더복잡하다.
// 프로세스 하나라도 쓰기를 한다면, 다른 어떤 프로세스도 그 동안에 접근할 수 없는 경우에 이용된다.

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
