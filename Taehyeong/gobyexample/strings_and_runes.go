package gobyexample

import (
	"fmt"
	"unicode/utf8"
)

// in other language
// string = character
// strings = character[]

// in Go
// string = byte[] = encodeUTF8(rune)
// strings = encodeUTF8(rune)[]

// rune 타입은 유니코드 포인트를 표현하는 int32
// rune 타입은 유니코드(UTF-8/Unicode)를 쉽게 제어하기 위한 타입
func stringsAndRunes() {
	// Thai 언어로 hello 라는 단어
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	// strings를 바이트의 배열로 출력
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// rune의 개수를 세기 위해 utf8 package 사용
	// UTF-8로 decode 되기 때문에 RuneCountInString은 string의 크기에 종속적
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

func examineRune(r rune) {
	// '' 안의 값은 rune의 값을 즉시 반환
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
