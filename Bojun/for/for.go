package main

import "fmt"

func main() {

	i := 1
	for i <= 3 { // Go doesn't have while
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { // for with init and post statement
		fmt.Println(j)
	}

	for { // infinite loop
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { //for with continue
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
