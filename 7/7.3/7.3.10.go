// 7.3.10 주의점
// 생산자 소비자 패턴

// 나쁜 예
c := make(chan int)
done := make(chan bool)
// 생산자 고루틴
go func() {
	for i := 0; i < 10; i++ {
		c <- i
	}
	done <- true
}()
// 소비자 고루틴 -> 문제점 : 끝나지 않음.
go func() {
	for {
		fmt.Println(<-c)
	}
}()
<-done