package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // int 선언 방식
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// int g
	// fmt.Println(g)

	f := "apple" // 선언 동시에 초기화
	fmt.Println(f)
}
