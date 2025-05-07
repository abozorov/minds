package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Owner   string
	Balance float64
}

// Выбрали поинтер ресивер из-за того, что нам нужно менять данные оригинала
func (a *Account) Deposit(amount float64) error {

	if amount <= 0 {
		return errors.New("Пополнение на число менеше или равное 0!")
	}

	a.Balance += amount
	fmt.Printf("Ваш счет был пополнен на %.2f сомони\n", amount)
	return nil
}

// Для обработки ошибки можно не bool использовать, а error как в функции Deposit
// Выбрали поинтер ресивер из-за того, что нам нужно менять данные оригинала
func (a *Account) Withdraw(amount float64) bool {

	if a.Balance-amount > 0 {
		a.Balance -= amount
		fmt.Printf("Снятие со счета %.2f сомони\n", amount)
		return true
	} else {
		return false
	}
}

// Простой вывод баланса поэтому валюе ресивер
func (a Account) GetBalance() float64 {
	return a.Balance
}

func main() {
	acc := Account{
		Owner:   "Kate",
		Balance: 0,
	}
	fmt.Printf("Добро пожаловать %s!\n Рады Вас видеть :)\n\n", acc.Owner)
	fmt.Println("Список доступных операций:")
	commands := " add \"ammount\" - попосление счета\n" + " wid \"ammount\" - снятие со счета\n" + " bal - выводит текущий баланс счета\n\n"
	fmt.Printf(commands)

	for {
		fmt.Println("Команда: ")
		var command string
		fmt.Scan(&command)

		switch command {
		case "add":
			fmt.Println("Сумма: ")
			var money float64
			fmt.Scan(&money)

			err := acc.Deposit(money)
			if err != nil {
				fmt.Println(err)
			}
		case "wid":
			fmt.Println("Сумма: ")
			var money float64
			fmt.Scan(&money)

			if !acc.Withdraw(money) {
				fmt.Printf("Недостаточно средств на счету!\nПопробуйте сумму поменьше!\n")
			}
		case "bal":
			fmt.Printf("Владелец счета %s, текущий баланс: %.2f\n", acc.Owner, acc.GetBalance())
		default:
			fmt.Printf("Неверная комманда. Попробуйте один из команд ниже\n%s", commands)
		}
	}
}
