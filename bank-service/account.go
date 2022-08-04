package bankservice

import (
	"errors"
	"fmt"
)

type Account struct {
	name string
	balance int
	password string
}

var insufficientBalance = errors.New("balance is insufficient")
var notExistAccount = errors.New("account not exist")
var accounts map[string]*Account

// 계좌 생성
func OpenAccount(name string, accountNumber string, password string) (string) {
	if accounts == nil{
		accounts = make(map[string]*Account)
	}

	_, exist := accounts[accountNumber] 
	
	if !exist{
		accounts[accountNumber] = &Account{name: name, balance: 0, password: password}
	}

	return accountNumber
}

// 입금
func Deposit(accountNumber string, amount int) error {
	_, exist := accounts[accountNumber]
	// 계좌 존재여부
	if !exist {
		return notExistAccount
	}

	accounts[accountNumber].balance += amount
	return nil
}

// 출금
func Withdraw(accountNumber string, amount int) error {
	_, exist := accounts[accountNumber]
	// 계좌 존재여부
	if !exist {
		return notExistAccount
	}
	// 잔액부족
	if accounts[accountNumber].balance < amount {
		return insufficientBalance

	}
	accounts[accountNumber].balance -= amount
	return nil
}

// 송금
func Transfer(senderAccountNumber string, payeeAccountNumber string, amount int) error {
	_, senderExist := accounts[senderAccountNumber]
	_, payeeExist := accounts[payeeAccountNumber]
	// 송금계좌 or 수취계좌가 존재 안할 경우
	if !senderExist || !payeeExist {
		return notExistAccount
	}
	// 송금자 잔액부족
	if accounts[senderAccountNumber].balance < amount {
		return insufficientBalance
	}
	accounts[senderAccountNumber].balance -= amount
	accounts[payeeAccountNumber].balance += amount
	return nil
}

// 잔액조회
func (a Account) getBalance() int {
	return a.balance
}

// 계좌이름 조회
func (a Account) getName() string {
	return a.name
}

// 특정 계좌 조회
func GetAccount(accountNumber ...string) string {
	var result string =""
	for _, v := range accountNumber {
		result = fmt.Sprintln(result, accounts[v].getName(), "'s account.\nAccountNumber: ", v, "\nHas balance: ", accounts[v].getBalance())
	}
	return result
}

// 전체 계좌 조회
func GetTotalAccount() string {
	var result string =""
	for key, val := range accounts {
		result = fmt.Sprintln(result, val.getName(), "'s account.\nAccountNumber: ", key, "\nHas balance: ", val.getBalance())
	}
	return result
}