package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите число для таблицы умножения")
	fmt.Scan(&number)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", i, number, number*i)
	}
}