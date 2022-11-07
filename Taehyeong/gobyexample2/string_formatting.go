package gobyexample2

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func stringFormatting() {

	p := point{1, 2}
	// {1 2}
	fmt.Printf("struct1: %v\n", p)
	// {x:1 y:2}
	fmt.Printf("struct2: %+v\n", p)
	// 소스코드 스니펫
	// main.point{x:1, y:2}
	fmt.Printf("struct3: %#v\n", p)
	// 값의 타입
	// main.point
	fmt.Printf("type: %T\n", p)
	// true
	fmt.Printf("bool: %t\n", true)
	// 123
	fmt.Printf("int: %d\n", 123)
	// 바이너리
	// 1110
	fmt.Printf("bin: %b\n", 14)
	// 정수에 해당하는 문자
	// !
	fmt.Printf("char: %c\n", 33)
	// 1c8
	fmt.Printf("hex: %x\n", 456)
	// 78.900000
	fmt.Printf("float1: %f\n", 78.9)
	// 1.234000e+08
	fmt.Printf("float2: %e\n", 123400000.0)
	// 1.234000E+08
	fmt.Printf("float3: %E\n", 123400000.0)
	// "string"
	fmt.Printf("str1: %s\n", "\"string\"")
	// 쌍 따옴표로 묶기
	// "\"string\""
	fmt.Printf("str2: %q\n", "\"string\"")
	// 6865782074686973
	fmt.Printf("str3: %x\n", "hex this")
	// 0x42135100
	fmt.Printf("pointer: %p\n", &p)
	// 너비 지정은 % 사용
	// |    12|   345|
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	// |  1.20|  3.45|
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	// |1.20  |3.45  |
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// |   foo|     b|
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	// |foo   |b     |
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	// 출력x 포매팅후 반환
	// a string
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	// an error
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
