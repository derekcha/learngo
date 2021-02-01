package accounts

// Account Struct
type account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}
