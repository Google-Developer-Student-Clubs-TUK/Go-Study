package gobyexample

import "fmt"

// go에서 array보다 일반적인 방식
// 배열 크기를 동적으로 변경가능, 부분 배열 발췌가능
func slices() {

	// make를 사용한 slice 선언
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// set
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	// get
	fmt.Println("get:", s[2])

	// slice의 길이
	fmt.Println("len:", len(s))

	// append를 이용한 확장
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy를 이용해 slice 복사가능
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// index를 지정해 배열 발췌가능
	// [low:high] low <= index < high
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// 또다른 선언법들
	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
	a[1] = 10
	fmt.Println(a) // [1, 10, 3]출력
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 다중 slice도 구성가능
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
