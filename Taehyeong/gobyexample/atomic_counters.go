package gobyexample

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicCounters() {
	// 부호없는 정수 사용을 위한 정의
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// non-atomic을 사용했다면 goroutine이 서로 간섭할 것이기 때문에
	// 실행 사이에 변경되는 다른 숫자를 얻을 것이다
	fmt.Println("ops:", ops)
}
