// 7.3.5 select
// select를 쓰면 여러 채널과 통신할 수 있다. ( switch와 비슷함 )

// 팬인하기
select {
case n := <-c1:
	c <- n
case n := <-c2:
	c <- n
case n := <-c3:
	c <- n
}