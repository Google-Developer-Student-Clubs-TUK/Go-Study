package gobyexample

import "fmt"

// 가변함수: 인자 갯수에 상관없이 호출가능
func sum(nums ...int) {
	// fmt.Println 도 가변함수
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func variadicFunctions() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// 슬라이스에 들어 있다면 func(slice...)를 사용
	sum(nums...)
}
