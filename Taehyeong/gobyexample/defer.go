package gobyexample

import (
	"fmt"
	"os"
)

// 한국말로로 지연함수
// 코드 정리 목적
func deferFuntion() {
	f := createFile("/tmp/defer.txt")
	// 함수 호출 보장
	// 함수가 끝나는 지점에서 실행
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
