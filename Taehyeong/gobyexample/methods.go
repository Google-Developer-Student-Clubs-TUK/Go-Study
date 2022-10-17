package gobyexample

import "fmt"

type rect struct {
	width, height int
}

// func 키워드와 함수명 사이 reciever 정의
// reciever: (method를 가지게될) struct 변수명 + struct 타입

// pointer receiver
// 데이터의 주소 전달, 호출자의 데이터 변경 o
func (r *rect) area() int {
	return r.width * r.height
}

// value receiver
// 데이터를 복사해 전달, 호출자의 데이터 변경 x
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func methods() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
