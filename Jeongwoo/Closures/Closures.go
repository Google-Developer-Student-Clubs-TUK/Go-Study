package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt()) //함수가 계속 실행될때마다 1더함
	fmt.Println(nextInt())

	newInts := intSeq() // 다시 지정하면 초기화됨. 폐쇄함수
	fmt.Println(newInts())
}
