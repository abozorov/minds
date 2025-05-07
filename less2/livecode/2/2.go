package main

import "fmt"

func main() {
	var weekDay string

	fmt.Scan(&weekDay)

	switch weekDay {
	case "понедельник":
		fmt.Println("Первый день недели")
	case "вторник":
		fmt.Println("Второй день недели")
	case "среда":
		fmt.Println("третий день недели")
	case "четверг":
		fmt.Println("четверный день недели")
	case "пятница":
		fmt.Println("пятый день недели")
	case "суббота":
		fmt.Println("шестой день недели")
	case "воскресенье":
		fmt.Println("седьмой день недели")
	default:
		fmt.Println("Ввели неверно")
	}
}
