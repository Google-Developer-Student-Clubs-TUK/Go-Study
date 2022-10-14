package main

import "fmt"

func zeroval(ival int) { // 값 복사
	ival = 0
}

func zeroptr(iptr *int) { // 포인터 넘김
	*iptr = 0 // 메모리 자체에 0을 저장
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i) // i자체의 주소
}
