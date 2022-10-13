package main

import "fmt"

func zeroval(ival int) {
	ival = 0
	fmt.Println("zeroval:", &ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
	fmt.Println("zeroptr:", iptr)
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) // call by value
	fmt.Println("zeroval:", i)

	zeroptr(&i) // call by refernce
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
