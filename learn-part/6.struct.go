package learnPart

import "fmt"

// type 정의
type hotel struct {
	name string
	price int
}

func StructFunc() {
	// 객체 생성
	var h1 = hotel{}
	// 필드값 설정
	h1.name = "name"
	h1.price = 100
	fmt.Println(h1)

	// struct 객체 생성 - Dereference (생략된 필드가 존재하면 Zero value로 자동 세팅된다.)
	h2 := hotel{name: "hotel2", price: 200}
	h3 := hotel{"hotel2", 200}
	// new: 모든 필드를 Zero value로 초기화 후 *hotel 리턴
	h4 := new(hotel)
	h4.name = "jayden"
	fmt.Println(h2, h3, h4)
}


type dict struct {
    data map[int]string
}
 
//생성자 함수 정의
func newDict() *dict {
    d := dict{}
    d.data = map[int]string{}
    return &d //포인터 전달
}

func ConstructorFunc()  {
	// struct 객체 생성 - reference
	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
}