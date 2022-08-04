package learnPart

func GotoFunc()  {
	num := 1
	if num == 1 {
		goto ONE
	} else {
		goto OTHER
	}

	ONE:
	println("num is 1")
	goto END
	OTHER:
	println("num is not 1")
	END:
}