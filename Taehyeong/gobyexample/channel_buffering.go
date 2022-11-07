package gobyexample

import "fmt"

// 기본상태는 버퍼x
// 제한된 버퍼 생성 가능
func channelBuffering() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
