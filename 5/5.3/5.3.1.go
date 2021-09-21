// 인터페이스 구조
interface {
	Method1()
	Method2(i int) error
}

//인터페이스에 이름을 붙일 수 있음
type Loader interface {
	Load(filename string) error
}

//여러 인터페이스 합칠 수 있음
type ReadWriter {
	io.Reader
	io.Writer
}