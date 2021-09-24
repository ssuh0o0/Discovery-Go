// 7.4.1 동시성 디버그

$ go test - race mypkg 		// 패키지 테스트
$ go run -race mysrc.go 	// 소스파일 run
$ go build -race mycmd 		// command를 build
$ go install -race mypkg	// pakage를 install

runtime.NumGoroutine()	// 현재 동작하는 고루틴의 수
runtime.NumCPU()		// 현재 사용 가능한 CPU 수
runtime.GOMAXPROCS()	// 얼마나 많은 CPU를 이용할 것인지