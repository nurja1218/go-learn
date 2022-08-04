package learnPart

func ConditionalFunc() {

	// if - 조건절 안에 문장을 넣을 수 있다
	i := 1
	if val := i * 2; val < 1 {
		println(val)
	}

	// switch 뒤에 조건변수 혹은 Expression을 적지 않는 경우
	num := 5
	switch {
	case num == 10:
		println("num is 10")
	case num == 5:
		println("num is 5")
	default:
		println("num is less than 5")
	}

	// 복수의 case 지정
	switch num{
	case 1, 2:
		println("num is 1 or 2")
	case 3, 4:
		println("num is 3 or 4")
	default:
		println("num is more than 5")
	}

	// fallthrough: case를 만족해도 아래의 case 들을 실행
	switch {
	case num > 8:
		println("A rank")
		fallthrough
	case num > 6:
		println("B rank")
		fallthrough
	case num > 4:
		println("C rank")
		fallthrough
	case num > 2:
		println("D rank")
		fallthrough
	default:
		println("long number")
	}


	// Expression을 사용한 경우
	var category = 1
	switch x := category << 2; x - 1 {
		//...
	}

	// // 타입검사
	// switch v.(type) {
	// case int:
	// 	println("int")
	// case bool:
	// 	println("bool")
	// case string:
	// 	println("string")
	// default:
	// 	println("unknown")
	// }

}