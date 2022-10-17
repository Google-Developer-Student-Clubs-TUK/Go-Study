package gobyexample

import "fmt"

// 키(Key)에 대응하는 값(Value)을 신속히 찾는 해시테이블(Hash table)
func maps() {
	// 선언방식1 make(map[key-type]val-type).
	m := make(map[string]int)

	// set and update
	m["k1"] = 7
	m["k2"] = 13

	// get
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 맵길이
	fmt.Println("len:", len(m))

	// key에 대응하는 key/value 쌍 삭제
	delete(m, "k2")
	fmt.Println("map:", m)

	// 변수를 읽을때 첫번재는 키의 값, 두번째는 존재 여부에 따른 bool 값
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 선언방식2 literal
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
