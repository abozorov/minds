package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число для проверки:")
	fmt.Scan(&num)

	fmt.Println("Число делиться на 3 и на 5?")
	if num%3 == 0 && num%5 == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
