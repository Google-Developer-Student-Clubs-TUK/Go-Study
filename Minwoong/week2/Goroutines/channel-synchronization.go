package main

import (
	"fmt"
	"time"
)

func workerB(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go workerB(done)

	<-done

}
