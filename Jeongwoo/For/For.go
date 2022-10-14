package main

import "fmt"

func main() {

	i := 1 // 선언하는 동시에 값 부여
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for { // 무한 반복
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { // n이 0부터 5까지
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
