package learnPart

import "fmt"

func ArrayFunc() {
	var a [3]int  //정수형 3개 요소를 갖는 배열 a 선언
    a[0] = 1
    a[1] = 2
    a[2] = 3
	
	// a1 = a2
	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3} //배열크기 자동으로
	println(a[1], a1[1], a2[1]) // 2 출력

	  // 정의
	  var ma = [2][3]int{
        {1, 2, 3},
        {4, 5, 6},  //끝에 콤마 추가
    }
    println(ma[1][2])

	var multiArray [3][4][5]int
	multiArray[0][1][2] = 10     // 사용
}

func SliceFunc()  {
	var s []int        //슬라이스 변수 선언 = 길이와 용량 0 = nil
    s = []int{1, 2, 3, 4, 5} //슬라이스에 리터럴값 지정
    s[1] = 10
    fmt.Println(s)     // [1, 10, 3]출력

	// but, 슬라이스 변수 선언 시 최대용량을 정의해야하는 make는 거의 쓰지 않는다
	s1 := make([]int, 5, 10)
    println(len(s1), cap(s1)) // len 5, cap 10

	// 부분 슬라이스
    s = s[2:5]  // = s[2:]
    fmt.Println(s) //3,4,5 출력

	// 슬라이스 추가, 병합, 복사
	s = append(s, 6, 7)
	fmt.Println(s) //3,4,5,6,7 출력

	s2 := []int{8, 9, 10}
	s = append(s, s2...)
	fmt.Println(s) //3,4,5,6,7,8,9,10 출력
}

func MapFunc()  {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	m["three"] = 3

	_, exists := m["three"]
    if !exists {
        println("empty three key")
    }

	for key, val := range m {
		println(key, val)
	}
}