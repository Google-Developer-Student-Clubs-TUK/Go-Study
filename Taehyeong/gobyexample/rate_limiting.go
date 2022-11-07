package gobyexample

import (
	"fmt"
	"time"
)

func rateLimiting() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 속도 제한기
	limiter := time.Tick(200 * time.Millisecond)

	// 매 요청이 바로 실행되지 않도록 조절
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 속도 제한을 유지하면서 버스팅
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
