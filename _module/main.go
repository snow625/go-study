package main

import (
	bank_account "bank_account/bankAccount"
	"fmt"
)

func main() {
	account := bank_account.CreateAccount("Dory")
	fmt.Println(account)
	account.DecreaseBalance(12)
	account.IncreaseBalance(100)
	balance := account.GetBalance()
	fmt.Println(balance)
}
