package learnPart

func DeferFunc()  {
	defer println("last")
	println("first")
}