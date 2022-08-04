package learnPart

import "fmt"

// type 정의
type hotel struct {
	name string
	price int
}

func DefinedTypeFunc() {

	var h1 = hotel{}
	h1.name = "name"
	h1.price = 100

	h2 := hotel{name: "hotel2", price: 200}

	fmt.Println(h1, h2)
}