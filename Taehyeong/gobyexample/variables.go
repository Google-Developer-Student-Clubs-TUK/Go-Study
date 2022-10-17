package gobyexample

import "fmt"

func variables() {
	// var: 타입 추정 가능
	var a = "initial"
	fmt.Println(a)

	// 한번에 여러 변수 선언
	// 변수 뒤 타입 지정 가능
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 초기화 되지 않은 변수 (zero-valued)
	// int, float = 0
	var d int
	fmt.Println(d)
	// string = ""
	var e string
	fmt.Println(e)
	// bool = false
	var f bool
	fmt.Println(f)

	// := 는 선언 연산자
	// 함수내에서만 사용가능
	// var g string = "apple"
	g := "apple"
	fmt.Println(g)
}
