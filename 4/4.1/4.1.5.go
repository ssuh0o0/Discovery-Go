// 가변 인수
// 받는 인자의 개수를 정하지 않은 함수 짜기

// 1번째 방법
func f(w io.writer, nums []int){
	...
}

// 2번째 방법
// ...을 써서 lines를 가변인자로 만들어줌 -> 그래도 슬라이스가 됨.
func WriteTo(w io.writer, lines ...string) (n int64, err error){

}

func main() {
	...
	// 1번 호출
	f(w, []int{x,y,z})

	// 2번 호출 - 1
	WriteTo(w, "hello", "world", "Go lang")

	// 2번 호출 -2
	lines := []string("hello", "world", "Go lang")
	WriteTo(w, lines...)
}