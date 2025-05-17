package main

import "fmt"

type BankAccount interface {
	Deposit(amount float64)
}

type SavingsAccount struct {
	Balance float64  
}

func (sa *SavingsAccount) Deposit(amount float64) {
	if amount < 0 {
		fmt.Println("Ошибка!")
		return
	}
	sa.Balance += amount
	fmt.Println("Успешно!")
}

type CheckingAccount struct {
	Account *SavingsAccount
}

func (sa *CheckingAccount) Deposit(amount float64) {

	if sa.Account.Balance > amount {
		fmt.Printf("Баланс больше чем %.2f!\n", amount)
	} else {
		fmt.Printf("Баланс меньше или равно чем %.2f!\n", amount)
	}
}

func someFunc(acc BankAccount, amount float64) {

	acc.Deposit(amount)
}

func main() {
	// Банковские счета
	//  - Интерфейс BankAccount с методами Deposit(amount float64),
	// Withdraw(amount float64), Balance() float64.
	//  - Структуры: SavingsAccount, CheckingAccount.
	//  - Разная логика по процентам или лимитам при снятии средств.
	var acc BankAccount
	acc = &SavingsAccount{
		Balance:0.0,
	}
	someFunc(acc, 124.1)
	someFunc(acc, -124.1)

	acc = &CheckingAccount{
		Account: acc.(*SavingsAccount),
	}
	someFunc(acc, 100)
	someFunc(acc, 200)
}
