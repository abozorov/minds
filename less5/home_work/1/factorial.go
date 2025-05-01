package main

import "fmt"

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// 1. Напишите рекурсивную функцию, которая вычисляет факториал числа.

	fmt.Println("Факториал какого натурального числа найти?")
	var n int64
	fmt.Scan(&n)

	fmt.Printf("Факториал числа %d равна %d\n", n, factorial(n))
}