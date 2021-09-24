// 7.2.4 닫힌 채널
package main

import "fmt"

// 채널이 열려있다면, 채널이 받을 값이 없을 때는 받을 값이 생길 때 까지 기다리지만,
// 채널이 닫혀있다면, 기다리지 않는다.
// 닫은 채널을 또 닫으면, 패닉이 발생한다.

func Example_ClosedChannel() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//output:
	//0
	//0
	//0

}
