//7.3.8 요청과 응답 짝짓기
// 분산처리되면 뭐가 값이 먼저 나올지 알 수 없다. 어느 요청에 의한 건지 알 필요가 없는 경우가 더 많지만, 알아야 하는 경우도 있다.
// ID번호도 같이 넘겨서 확인해보는 방법이 있ㄷ다.
// 또, 보내는 쪽에서 요청 ID를 보관하고 있지만, 요청에 대한 응답을 다른 고루틴이 받아갈 수 있다면 이것도 골치아프다.
// 내가 보낸 요청에 대한 응답을 내가 받기 위해서는 테크닉이 필요하다.

package main

import (
	"fmt"
	"sync"
)

type Request struct {
	Num  int
	Resp chan Response
}
type Response struct {
	Num      int
	WorkerID int
}

func PlusOneService(reqs <-chan Request, workerID int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerID}
		}(req)
	}
}

func main() {
	reqs := make(chan Request)
	defer close(reqs)
	for i := 0; i < 3; i++ {
		go PlusOneService(reqs, i)
	}
	var wg sync.WaitGroup
	for i := 3; i < 53; i += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resps := make(chan Response)
			reqs <- Request{i, resps}
			fmt.Println(i, "=>", <-resps)
			// 여러개의 결과를 받아야 하는 경우에는 이렇게 하자.
			// for resp := range resps {
			// 	fmt.Println(i, "=>", resp)
			// }
		}(i)
	}
	wg.Wait()
}
