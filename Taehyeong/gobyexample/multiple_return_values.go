package gobyexample

import "fmt"

func vals() (int, int) {
	// multiple return values
	return 3, 7
}

func multipleReturnValues() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 두번째 반환값만 받기위해선 _ 처리
	_, c := vals()
	fmt.Println(c)
}
