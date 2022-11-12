package gobyexample

import (
	"fmt"
	"time"
)

// 고루틴간 채널 동기화
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSynchronization() {
	done := make(chan bool, 1)
	go worker(done)

	// 채널은 worker 함수로부터 알림을 받을 때까지 블로킹
	<-done
}
