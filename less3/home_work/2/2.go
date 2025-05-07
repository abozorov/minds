package main

import "fmt"

func main() {

	// 	Удалить элемент из среза по индексу.
	//  - Дан срез из 5 элементов, удалить, например, третий элемент.
	//  > Подсказка: slice = append(slice[:index], slice[index+1:]...)

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var index int
	fmt.Println("Какой элемент удалить?")
	fmt.Scan(&index)
	a = append(a[:index-1], a[index:]...)
	fmt.Println("Массив после удаления элемента:", a)
}
