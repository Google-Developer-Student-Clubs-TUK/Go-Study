package gobyexample

import "fmt"

func mayPanic() {
	panic("a problem")
}

func recoverFunction() {

	// defer을 이용한 panic 포착
	defer func() {
		if r := recover(); r != nil {
			// 반환값은 panic
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
