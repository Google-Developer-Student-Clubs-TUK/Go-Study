package gobyexample

import (
	"fmt"
	"math"
)

// 변하지않는 상수 const
const s string = "constant"

func constants() {
	fmt.Println(s)

	const n = 500000000

	// arbitrary precision arithmetic 방식으로 큰 수 표현
	const d = 3e20 / n
	fmt.Println(d)

	// 형변환
	fmt.Println(int64(d))

	// 함수에 할당 가능
	fmt.Println(math.Sin(n))

	// 한번에 여러 변수 선언 가능
	const (
		x = "x"
		y = "y"
		z = "z"
	)
	fmt.Println(x, y, z)

	// iota를 이용하면 상수값을 0부터 선언가능
	const (
		q = iota // 0
		w        // 1
		e        // 2
	)
	fmt.Println(q, w, e)
}
