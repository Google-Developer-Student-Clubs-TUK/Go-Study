package gobyexample

import "fmt"

func foFunction() {

	// single condition 방식
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// initial/condition/after for loop 방식
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 조건없는 for문은 break 또는 함수 return 전까지 반복
	for {
		fmt.Println("loop")
		break
	}

	// continue를 이용한 다음 루프 진행
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
