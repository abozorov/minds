package main

import "fmt"

func someFunc(n int, f func(n int) int) int {
	return f(n)
}

func main() {
	// 	2. Напишите функцию, которая принимает число и функцию, и применяет её к
	// этому числу.
	
	fmt.Println("Введите число")
	var n int
	fmt.Scan(&n)

	fmt.Printf("Увеличение числа %d на единицу: %d\n", n, someFunc(n, func(n int) int { // тут можно реализовать любую функцию.
		return n+1
	}))

	fmt.Printf("Возведение числа %d в квадрат: %d\n", n, someFunc(n, func(n int) int {
		return n*n
	}))

	fmt.Printf("Умножение числа %d на 2: %d\n", n, someFunc(n, func(n int) int {
		return n*2
	}))
}
