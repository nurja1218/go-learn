package learnPart

func LoopFunc()  {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	println("sum: ", sum)

	i, sum2 := 0, 0
	for i < 10 {
		i++
		sum2 += i
	}
	println("sum2: ", sum2)

	// range loop
	numbers := []int{10, 11, 12}

	for index, num := range numbers {
		println(index, num)
	}

	// break
	i = 0
	for i < 10 {
		i++
		if i == 5 {
			break
		}
	}
	println("i :", i)

	// continue
	i = 0
	for i < 5 {
		i++
		if i == 3 {
			continue
		}
		println("i :", i)
	}
}