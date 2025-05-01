package main

import "fmt"

func isPrime(n int) bool {

	for i := 2; i*i <= n; i++ {

		if n%i == 0 {
			return false
		}
	}
	return true && n != 1
}

func main() {
	// 3. Напишите функцию, которая проверяет, является ли число простым.

	fmt.Println("Введите число")
	var n int
	fmt.Scan(&n)

	fmt.Printf("Число %d простое?\n %t", n, isPrime(n))
}
