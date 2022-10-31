package gobyexample2

import "fmt"

func nonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan bool)

	// msg 가 없으므로 바로 default 리턴
	// 채널을 수신하기전까지 블로킹하는것을 방지
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
