package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	account "github.com/nurja1218/go-learn/bank-service"
	"github.com/nurja1218/go-learn/logis"
)

const (
	host     = "127.0.0.1"
	database = "logis"
	user     = "jayden"
	password = "jayden1"
)

func main() {
	var conn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, password, host, database)

	db, err := sql.Open("mysql", conn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// // user 조회
	// users := logis.GetAllUser(db)
	// for _, user := range users {
	// 	fmt.Printf("%v\n", user.Name)
	// }

	// // device 조회
	// devices := logis.GetAllDevice(db)

	// for _, device := range devices {
	// 	fmt.Printf("%v\n", device.Serial)
	// }

	// company 조회
	companies := logis.GetAllCompany(db)

	for _, company := range companies {
		fmt.Printf("%v\n", company.Name)
	}
}

func callBankService() {
	jaydenAccountNum := account.OpenAccount("jayden", "1234", "1234")
	ziniAccountNum := account.OpenAccount("zini", "1111", "1234")
	_ = account.OpenAccount("evie", "2222", "1234")

	// 1. jayden 10만원 입금
	error1 := account.Deposit(jaydenAccountNum, 100000)
	if error1 != nil {
		fmt.Println("error message: ", error1)
	}
	// 2. zini 30만원 출금
	error2 := account.Withdraw(ziniAccountNum, 300000)
	if error2 != nil {
		fmt.Println("error message: ", error2)
	}
	// 3. evie 계좌 조회
	//  evieAccount := account.GetAccount(evieAccountNum)
	//  println(evieAccount)
	// 4. jayden -> zini 5만원 송금
	error3 := account.Transfer(jaydenAccountNum, ziniAccountNum, 30000)
	if error3 != nil {
		fmt.Println("error message: ", error3)
	}
	// 5. 전체 계좌 조회
	// result := account.GetAccount(evieAccountNum, jaydenAccountNum, ziniAccountNum)
	result := account.GetTotalAccount()
	println(result)
}
