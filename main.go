package main

import (
	"fmt"

	"github.com/derekcha/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Derek")
	account.Deposit(30)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
