package main

import "fmt"

func main() {
	//  - Создать map с переводами слов (английский → русский).
	//  - Реализовать поиск перевода словаб.

	fmt.Println("\nПривет, это программа \"консольный словарь\".")
	fmt.Println("\nВам доступны 2 команды:\n ")
	fmt.Println(" 1. add \"слово\" \"перевод\" - которая ДОБАВЛЯЕТ слово в словарь")
	fmt.Println(" 2. find \"слово\" - которая ИЩЕТ слово в словаре")
	fmt.Println("\nДля выхода из программы введите команды \"exit\"")

	dictionary := make(map[string]string)

	// Перевод что на английском что на русском. В мапу слова добавлены в обе стороны
	for {
		var command string
		fmt.Scan(&command)
		if command == "add" {
			var word, translate string
			fmt.Scan(&word, &translate)

			dictionary[word] = translate
			dictionary[translate] = word
		} else if command == "find" {
			var word string
			fmt.Scan(&word)
			translate, found := dictionary[word]

			if found {
				fmt.Printf(" Перевод слова %s - %s\n\n", word, translate)
			} else {
				fmt.Printf(" В словаре нет слова %s, вы можете его добавить с помошью команды add\n\n", word)
			}
		} else if command == "exit" {
			break
		} else {
			fmt.Println(" Введена неверная команда, пожалуйста повторите запрос!\n ")
			fmt.Scanln()
		}
	}
}

/*
add apple яблоко
add space космос
add green зеленый
add красный red
add yellow желтый
add desk доска
add black черный
add логотип logo

find logo
find apple
find зеленый
fing 90
find hello
find space
*/
