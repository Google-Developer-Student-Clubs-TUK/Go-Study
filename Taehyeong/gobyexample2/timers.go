package gobyexample2

import (
	"fmt"
	"time"
)

func timers() {
	// 2초후 알림
	timer1 := time.NewTimer(2 * time.Second)

	// 타이머 만료전 타이머 블로킹
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 단순한 대기용, 만료전 취소 불가능
	time.Sleep(2 * time.Second)
}
