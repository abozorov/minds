package main

import "fmt"

func main() {

	// 	Посчитать количество чётных чисел в массиве.
	//  - Пройти по всем элементам и подсчитать.

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	even := 0

	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			even++
		}
	}
	fmt.Printf("even = %d\n", even)

}
