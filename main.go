package main

import (
	account "github.com/nurja1218/golearn/bank-service"
)

func main()  {
	 jaydenAccountNum := account.OpenAccount("jayden", "1234", "1234")
	 ziniAccountNum := account.OpenAccount("zini", "1111", "1234")
	 evieAccountNum := account.OpenAccount("evie", "2222", "1234")

	 // 1. jayden 10만원 입금
	 account.Deposit(jaydenAccountNum, 100000)
	 // 2. zini 30만원 출금
	 account.Withdraw(ziniAccountNum, 300000)
	 // 3. evie 계좌 조회
	 evieAccount := account.GetAccount(evieAccountNum)
	 println(evieAccount)
	 // 4. jayden -> zini 5만원 송금
	 account.Transfer(jaydenAccountNum, ziniAccountNum, 50000)
	 // 5. 전체 계좌 조회
	result := account.GetTotalAccount()
	println(result)
}