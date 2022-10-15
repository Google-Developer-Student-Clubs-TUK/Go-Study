package gobyexample

import "fmt"

// 고정된 배열크기 안에 동일한 타입의 데이타를 연속적으로 저장
func arrays() {
	// type의 초기값이 zero-valued 인 일정 사이즈의 배열 선언
	var a [5]int
	fmt.Println("emp:", a)

	// set
	a[4] = 100
	fmt.Println("set:", a)
	// get
	fmt.Println("get:", a[4])

	// array length
	fmt.Println("len:", len(a))

	// 배열 초기값 설정
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 다중배열 선언 가능
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 일반적인 다중배열 초기화 방식
	var e = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가
	}
	fmt.Println("a: ", e[1][2])
}
