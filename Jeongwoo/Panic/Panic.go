package main

import "os"

func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file") // 오류 메세지와 함께 정상적으로 출력 불가능
	if err != nil {
		panic(err)
	}
}
