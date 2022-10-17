package gobyexample

import "fmt"

func plus(a int, b int) int {
	// 반환값이 있다면 return
	return a + b
}

// 같은 타입을 갖는 파라미터들이 연속적이면 마지막빼고 생략가능
func plusPlus(a, b, c int) int {
	return a + b + c
}

func functions() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
