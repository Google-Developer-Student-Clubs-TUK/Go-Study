package gobyexample

import "os"

func panicFunction() {

	// 예상치 못한 에러를 체크
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		// 예상치 못한 에러값을 반환했을때 중단
		panic(err)
	}
}
