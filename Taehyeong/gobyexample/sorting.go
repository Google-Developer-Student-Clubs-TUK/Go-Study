package gobyexample

import (
	"fmt"
	"sort"
)

func sorting() {

	strs := []string{"c", "a", "b"}
	// 전달된 슬라이스를 변형하는 형태로 새로운 슬라이스 반환 x
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 정렬여부 검사
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
