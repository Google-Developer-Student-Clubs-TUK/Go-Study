package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true //boolean 형태
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done // 지우면 worker 실행되기 전에 종료될수도 있음
}
