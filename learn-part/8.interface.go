package learnPart

import "fmt"

// interface: 메서드들의 집합체
type reserve interface {
	twoday() int
}

type airbnb struct {
	name   string
	price  int
	coupon int
}

func (h hotel) twoday() int {
	return h.price * 2
}

func (a airbnb) twoday() int {
	return a.price*2 - a.coupon
}

func makeReserve(r ...reserve) {
	fmt.Println(len(r))
	fmt.Println(r)
	for _, v := range r {
		fmt.Println(v)
		fmt.Println(v.twoday())
	}
}

func InterfaceFunc() {
	a := hotel{"aaa", 1000}
	b := airbnb{"bbb", 1000, 500}

	makeReserve(a, b)
}

