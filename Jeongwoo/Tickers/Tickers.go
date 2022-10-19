package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool) // 500ms마다 계속 동작

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond) // 1600 후 동안 기다리고 중지
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
