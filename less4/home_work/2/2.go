package main

import "fmt"

func main() {
	// Написать программу, которая:
	// 	- Создает map[string]int, где ключ — название товара, а значение — его цена.
	// 	- Позволяет добавлять новые товары и узнавать цену по названию.
   
	fmt.Println("\nПривет, это программа \"консольный склад\".")
	fmt.Println("\nВам доступны 2 команды:\n ")
	fmt.Println(" 1. add \"товар\" \"цена\" - которая ДОБАВЛЯЕТ товар в список")
	fmt.Println(" 2. find \"название товара\" - которая ИЩЕТ товар в списке")
	fmt.Println("\nДля выхода из программы введите команды \"exit\"")

	storage := make(map[string]int)
	for {
		var command string
		fmt.Scan(&command)
		if command == "add" {
			var(
				productName string
				price int
			)
			fmt.Scan(&productName, &price)
			storage[productName] = price
		} else if command == "find" {
			var productName string
			fmt.Scan(&productName)
			price, found := storage[productName]

			if found {
				fmt.Printf(" Цена товара %s - %d СОМОНИ\n\n", productName, price)
			} else {
				fmt.Printf(" На складе нет товара с наименованием %s, вы можете его добавить с помошью команды add\n\n", productName)
			}
		} else if command == "exit" {
			break
		} else {
			fmt.Println(" Введена неверная команда, пожалуйста повторите запрос!\n ")
		}
	}
}

/*
add pencil 1
add laptop 5120
add notebook 5
add iphone16 13450
add desk 354
add tv 4500
add mouse 60
add keyboard 220

find mouse
find laptop
fing d
find pencil
find f
*/
