package learnPart

func ArrayFunc() {
	
}

func SliceFunc()  {
	
}

func MapFunc()  {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	m["three"] = 3

	for k, v := range m {
		println(k, v)
	}
}