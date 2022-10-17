package gobyexample

import "fmt"

func rangeFunction() {

	nums := []int{2, 3, 4}
	sum := 0
	// slice 순회
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 첫 인자는 index, 두번재 인자는 대응하는 원소
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// map에서의 key,value 쌍 순회
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 문자열에서 유니코드 코드 포인트 순회
	// 첫번째 값은 rune 시작 바이트 위치
	// 두번째 값은 rune 의 값
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
