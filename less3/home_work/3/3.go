package main

import "fmt"

func swap(a, b *int) {
	*a += *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	// 	Перевернуть срез.
	//  - Сделать так, чтобы элементы шли в обратном порядке.

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n/2; i++ {
		swap(&a[i], &a[n-i-1])
	}
	fmt.Println("Перевернутый массив:", a)
}
