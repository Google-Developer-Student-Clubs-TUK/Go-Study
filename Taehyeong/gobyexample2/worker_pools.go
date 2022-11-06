package gobyexample2

import (
	"fmt"
	"time"
)

func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// 여러 인스턴스 동시 실행
func workerPools() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// worker 3개를 보내지만 현재 잡이 없으므로 블로킹상태
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	// 잡을 보내 일을시킴
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
