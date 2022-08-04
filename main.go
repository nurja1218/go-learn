package main

import (
	"fmt"

	account "github.com/nurja1218/golearn/bank-service"
)

func main()  {
	 jaydenAccountNum := account.OpenAccount("jayden", "1234", "1234")
	 ziniAccountNum := account.OpenAccount("zini", "1111", "1234")
	 _ = account.OpenAccount("evie", "2222", "1234")

	 // 1. jayden 10만원 입금
	 error1 := account.Deposit(jaydenAccountNum, 100000)
	 if error1 != nil{
		fmt.Println("error message: ", error1)
	 }
	 // 2. zini 30만원 출금
	 error2 :=account.Withdraw(ziniAccountNum, 300000)
	 if error2 != nil{
		fmt.Println("error message: ", error2)
	 }
	 // 3. evie 계좌 조회
	//  evieAccount := account.GetAccount(evieAccountNum)
	//  println(evieAccount)
	 // 4. jayden -> zini 5만원 송금
	 error3 := account.Transfer(jaydenAccountNum, ziniAccountNum, 30000)
	 if error3 != nil{
		fmt.Println("error message: ", error3)
	 }
	 // 5. 전체 계좌 조회
	// result := account.GetAccount(evieAccountNum, jaydenAccountNum, ziniAccountNum)
	result := account.GetTotalAccount()
	println(result)
}