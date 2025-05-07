package main

import "fmt"

func nFibonacci(n int) (fibonacci []int) {
	fibonacci = make([]int, 0, n)

	if n == 0 {
		return
	}
	fibonacci = append(fibonacci, 0)

	if n == 1 {
		return
	}
	
	fibonacci = append(fibonacci, 1)
	if n == 2 {
		return
	}

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
