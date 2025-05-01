package main

import "fmt"

func mapInt(a []int, f func(n int) int) (cache []int) {
	cache = make([]int, 0, len(a))

	for _, v := range a {
		cache = append(cache,  f(v))
	}
	return
}

func main() {
	// 	5. (Дополнительно) Напишите функцию mapInt, которая принимает слайс чисел и
	// функцию, и возвращает новый слайс с результатами применения функции к
	// каждому элементу.

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Printf("Увеличение каждого элемента слайса на единицу.\n Сам слайс: %v\n Изменения: %v\n", a, mapInt(a, func(n int) int {
		return n+1
	}))

	fmt.Printf("Возведение каждого элемента слайса в квадрат.\n Сам слайс: %v\n Изменения: %v\n",  a, mapInt(a, func(n int) int {
		return n*n
	}))
}
