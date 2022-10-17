package gobyexample

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// 포인터 사용가능
func zeroptr(iptr *int) {
	// 주소로 역참조하여 값 초기화
	*iptr = 0
}

func pointers() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 주소값 전달
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
