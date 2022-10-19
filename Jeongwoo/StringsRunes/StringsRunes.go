package main

import (
	"fmt"
	"unicode/utf8" // 라이브러리 문자열 -> UTF-8 인코딩 텍스트 컨테이너
)

func main() {

	const s = "สวัสดี" // 문자열 UTF-8로 인코딩

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) { //룬 -> 문자열의 유니코드 코드 포인트

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// strings -> char의 배열, string -> byte배열 사용 (rune의 UTF-8로 인코딩)
// string의 배열 -> UTF-8로 인코딩된 룬을 배열로
// byte배열 상태로 가지고 있다가 print할때 룬으로 인코딩해줘서
// 기본적으로 Rune형태가 아니고 byte자체를 인코딩하면 rune이 된다.
// 영어는 1바이트, 다른 글자는 3바이트 (한글)
