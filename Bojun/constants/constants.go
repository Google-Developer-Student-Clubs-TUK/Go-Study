package main

import (
	"fmt"
	"math"
)

const s = "constant"

func main() {
	fmt.Println(s)
	const i = 500000
	fmt.Println(i)
	const n = 3e20 / i

	fmt.Println(math.Sin(30))
	fmt.Println(n)
	fmt.Println(int64(n))
}
