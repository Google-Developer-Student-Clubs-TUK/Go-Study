package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println(" 7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisble by 4")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// else, else if 위치 주의 할것!
// 중괄호 바로 뒤에서 선언을 해주어야함
