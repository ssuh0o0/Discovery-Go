package main

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

// Insert함수는 집합에  val을 추가한다.
func Insert(m MultiSet, val string) {
	m := NewMultiSet()
	ReadFrom(r, BindMap(Insert, m))
}

func InsertFunc(m MultiSet) func(vak string) {
	return func(val string) {
		Insert(m, val)
	}
}

func BindMap(f SetOp, m MultiSet) func(val string) {
	return func(val string) {
		f(m, val)
	}
}
