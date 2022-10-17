package gobyexample

import "fmt"

func values() {
	// string
	fmt.Println("go" + "lang")

	// int
	fmt.Println("1+1 =", 1+1)
	// floats
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// bool
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
