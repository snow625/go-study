package bank_account

import "fmt"

type Account struct {
	id      int
	name    string
	balance float64
}

func newAccount(name string) Account {
	if name == "" {
		return Account{}
	}
	return Account{
		name:    name,
		balance: 0.00,
		id:      23232323,
	}
}

func CreateAccount(name string) Account {
	fmt.Println("Creating Account")
	return newAccount(name)
}

func (a Account) GetName() string {
	return a.name
}

func (a *Account) UpdateName(newName string) {
	if newName != "" {
		a.name = newName
	}
}
