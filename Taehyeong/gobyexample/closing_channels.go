package gobyexample

import "fmt"

func closingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
				// more 값이 flase이면 close
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// close 실행시 more은 false
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
