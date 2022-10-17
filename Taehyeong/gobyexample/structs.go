package gobyexample

import "fmt"

// 필드 데이타만을 가지며, (행위를 표현하는) 메서드를 갖지 않는다
type person struct {
	name string
	age  int
}

// 생성자함수
// new person struct 를 만드는 함수
func newPerson(name string) *person {
	// 안전하게 지역변수에 포인터 리턴
	p := person{name: name}
	p.age = 42
	return &p
}

func structs() {

	fmt.Println(person{"Bob", 20})
	// 구조체 필드 이름 지정 가능
	fmt.Println(person{name: "Alice", age: 30})
	// 제외 된 필드의 값은 zero-value
	fmt.Println(person{name: "Fred"})
	// 앞에 &를 붙이면 구조체 포인터를 생성
	fmt.Println(&person{name: "Ann", age: 40})
	// 생성자 함수에서 새로운 구조체 생성을 캡슐화하는 것은 관용적
	fmt.Println(newPerson("Jon"))
	// .을 이용해 접근 가능
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// 구조체 포인터에 .(dot)을 사용하면 자동으로 역참조
	sp := &s
	fmt.Println(sp.age)
	// 구조체는 mutable 하다
	sp.age = 51
	fmt.Println(sp.age)
}
