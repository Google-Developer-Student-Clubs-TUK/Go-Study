package gobyexample

import "fmt"

// 재귀함수
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func recursion() {
	fmt.Println(fact(7))

	// 클로저로 재귀를 구현 가능
	// 선언과 동시에 사용할 수 없으므로 먼저 var 타입 클로저를 선언해야함
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
