package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue { // 2개의 범위 설정
		fmt.Println(elem)
	}
}
