package gobyexample

import "fmt"

// 비어있지 않은 채널을 닫아도 채널에 남아있는 값들은 수신할 수 있다
func rangeOverChannels() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
