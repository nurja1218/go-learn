package learnPart

func ChanFunc()  {
	ch := make(chan int)

	go func ()  {
		ch <- 1
	}()

	num := <- ch

	println(num)
}