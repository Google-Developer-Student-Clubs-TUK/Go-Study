package gobyexample

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// base를 embed 하고 있는 컨테이너
type container struct {
	base
	str string
}

// 타입의 결합을 위한 구조체와 인터페이스의 embedding 지원
func structEmbedding() {

	// embedding struct를 생성할때 embedding 을 명시적으로 초기화해야함
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// embedding 필드를 직접 접근가능
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// fullpath 사용 가능
	fmt.Println("also num:", co.base.num)

	// embedding method 사용 가능
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// 컨테이너는 embeds 되어 있는 인터페이스로 다른 인터페이스를 implements 할 수 있다
	var d describer = co
	fmt.Println("describer:", d.describe())
}
