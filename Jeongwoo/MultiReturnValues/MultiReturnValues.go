package main

import "fmt"

func vals() (int, int) {
	return 3, 7 // 한번에 불러오기 가능
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
