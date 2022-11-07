package gobyexample

import "fmt"

// 고루틴 사이 파이프
// 송신 수신이 블로킹으로 이루어지듯
// 다른 동기화 작업 필요 없이 프로그램 종료를 기다리게 만들 수 있음
func channels() {
	// 새로운 채널 생성
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
