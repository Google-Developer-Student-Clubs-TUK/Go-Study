package gobyexample

import (
	"fmt"
	"math"
)

// 타입(type)이 구현해야 하는 메서드 원형(prototype)들을 정의
type geometry interface {
	area2() float64
	perim2() float64
}

type rect2 struct {
	width, height float64
}
type circle2 struct {
	radius float64
}

// 인터페이스에 있는 모든 메서드들을 구현
func (r rect2) area2() float64 {
	return r.width * r.height
}
func (r rect2) perim2() float64 {
	return 2*r.width + 2*r.height
}

func (c circle2) area2() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle2) perim2() float64 {
	return 2 * math.Pi * c.radius
}

// 변수가 인터페이스 타입을 가지면 해당 인터페이스 내 메서드들을 호출 가능
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area2())
	fmt.Println(g.perim2())
}

func interfaces() {
	r := rect2{width: 3, height: 4}
	c := circle2{radius: 5}

	measure(r)
	measure(c)
}
