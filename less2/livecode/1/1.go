package main

import "fmt"

func main() {
	fmt.Println("Введите возраст:")
	var age int
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Совершеннолетний")
	} else {
		fmt.Println("Несовершеннолетний")
	}
}
