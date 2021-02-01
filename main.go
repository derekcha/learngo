package main

import (
	"fmt"

	"github.com/derekcha/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Derek")
	fmt.Println(account)
	account.Deposit(100)
	fmt.Println(account.Balance())
}
