package gobyexample

import "fmt"

// 송신용 채널
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 수신용 채널
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
