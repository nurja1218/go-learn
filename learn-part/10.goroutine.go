package learnPart

import "time"

// go - 논리적(가상) 쓰레드
func GoRoutineFunc() {
	num(11)
	go num(1)
	go num(2)
	go num(3)
	time.Sleep(time.Second*3)
}

func num(n int)  {
	println(n)
}