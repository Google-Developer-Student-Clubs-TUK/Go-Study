package gobyexample

import "fmt"

// Closure는 함수 바깥에 있는 변수를 참조하는 함수값
// 함수를 인자로도, 리턴도 가능
func intSeq() func() int {
	i := 0
	// 이름 없이 한줄로 함수를 정의 가능
	// 익명함수 자체가 로컬 변수로 i 를 가지지 않으므로 안의 함수만 갱신하면 i는 초기화x
	return func() int {
		i++
		return i
	}
}

func closures() {

	// nextInt를 호출할 때마다 i 값 업데이트
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
