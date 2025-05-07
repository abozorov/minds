package main

import "fmt"

func main() {
	var age int

	fmt.Scan(&age)

	switch {
	case age < 18:
		fmt.Println("Школота")
	case 18 <= age && age <= 25:
		fmt.Println("Студент")
	case 26 <= age && age <= 60:
		fmt.Println("Работяга")
	case age > 60:
		fmt.Println("Пенсионер")
	}
}
