package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	nums[0] = 100
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	b := []int{1, 2, 3, 4, 5}

	sum(b...)
	fmt.Println(b)
}
