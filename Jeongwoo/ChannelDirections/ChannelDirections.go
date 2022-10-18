package main

import "fmt"

func ping(pings chan<- string, msg string) { // 오직 보내는 역할
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // 오직 받는 역할
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
