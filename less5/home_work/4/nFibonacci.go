package main

import "fmt"

func nFibonacci(n int) (fibonacci []int) {
	fibonacci = make([]int, 0, n)
	fibonacci = append(fibonacci, 0, 1)

	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	return
}

func main() {
	// 4. Напишите функцию, которая возвращает первые n чисел Фибоначчи в слайсе.
	
	fmt.Println("Введите число")
	var n int
	fmt.Scan(&n)

	fmt.Printf("Первые %d чисел Фибоначчи %v", n, nFibonacci(n))
}
