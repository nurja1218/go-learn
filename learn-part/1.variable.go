package learnPart

func VariableFunc(){
	// 변수, 상수
	var a int
	a= 11
	var b, c int = 2, 3
	const d int = 1
	const (
		e = 9
		f = 8
	)
	println(a, b, c, d, e, f)

	var g float32 = 11.
	// iota: 0부터 순차적으로 1씩 증가하여 변수에 할당
	const (
		num1 = iota
		num2
		num3
	)
	println(g, num1, num2, num3)

	// 데이터 타입변환
	var i int = 100
    var u uint = uint(i)
    var ft float32 = float32(i)  
    println(ft, u)
 
    str := "ABC"
    bytes := []byte(str)
    str2 := string(bytes)
    println(bytes, str2)
	
}