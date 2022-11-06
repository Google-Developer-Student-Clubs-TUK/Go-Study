package gobyexample2

import (
	"fmt"
	"sync"
	"time"
)

func worker3(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func waitGroups() {
	// 모든 고루틴이 끝날때까지 기다리게함
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker3(i)
		}()
	}

	// 포인터를 쓰지않고도 끝날때까지 기다림
	wg.Wait()
}
