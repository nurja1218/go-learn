package learnPart

//Rect - struct 정의
type Rect1 struct {
    width, height int
}
type Rect2 struct {
    width, height int
}
 
//Value receiver - struct를 복사해와서 변수 저장
func (r Rect1) area() int {
    return r.width * r.height   
}

// Pointer Receiver - 해당 struct에 직접 접근하여 데이터 변경
func (r *Rect2) area2() int {
    r.width++
    return r.width * r.height
}
 
func MethodFunc() {
    rect1 := Rect1{10, 20}
    area1 := rect1.area() //메서드 호출
    println(area1)

	rect2 := Rect2{10, 20}
    area2 := rect2.area2() //메서드 호출
    println(rect2.width, area2) // 11 220 출력

}
