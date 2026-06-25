package bank_account

import "fmt"

func (a *Account) IncreaseBalance(amount float64) {
	fmt.Println("IncreaseBalance")
	if amount > 0 {
		a.balance = a.balance + amount
	}

}

func (a *Account) DecreaseBalance(amount float64) {
	fmt.Println("DecreaseBalance")
	if amount > 0 {
		a.balance = a.balance - amount
	}

}

func (a Account) GetBalance() float64 {
	return a.balance
}
