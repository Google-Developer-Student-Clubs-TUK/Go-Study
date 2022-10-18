package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }() // <- 채널 구문 (고루틴 연결)

	msg := <-messages
	fmt.Println(msg)
}
