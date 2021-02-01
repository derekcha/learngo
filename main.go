package main

import (
	"fmt"

	"github.com/derekcha/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Derek")
	fmt.Println(account)
}
