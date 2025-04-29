package main

import "fmt"

func main() {
	// Дан массив из N элементов(стринги):
	// 1. Вывести не повторящиеся элементы и длину
	// 2. Вывести кол-во вхождений каждой из элементов

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	us := make(map[string]int)
	strs := make([]string, n)

	fmt.Println("Введите элементы массива(Строки) через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&strs[i])
		us[strs[i]]++
	}

	unic := make([]string, 0, n)
	for i := 0; i < n; i++ {
		if us[strs[i]] == 1 {
			unic = append(unic, strs[i])
		}
	}
	fmt.Println( "Всего неповторщихся элементов:", len(unic), "Сами элементы:", unic)

	for key, val := range us {
		fmt.Printf("Элемент \"%s\" и кол-во вхождений %d\n", key, val)
	}
}
