package main

import "fmt"

func swap(a, b *int) {
	*a += *b
	*b = *a - *b
	*a = *a - *b
}

func mergeSort(a []int, l, r int) {

	if l+1 >= r {
		if a[l] > a[r] {
			swap(&a[l], &a[r])
		}
	} else {
		m := (l + r) / 2
		mergeSort(a, l, m)
		mergeSort(a, m+1, r)
		cach := make([]int, 0, r)
		ll, x := l, m+1

		for x <= r || l <= m {

			if l <= m && (x > r || a[l] < a[x]) {
				cach = append(cach, a[l])
				l++
			} else if x <= r {
				cach = append(cach, a[x])
				x++
			}
		}

		for i := 0; i < len(cach); i++ {
			a[i+ll] = cach[i]
		}
	}
}

func problem1() {
	fmt.Println("Условие задачи:\n Найти сумму всех элементов слайса.\n ")

	fmt.Println("Введите кол-во элементов:")
	var n int
	fmt.Scan(&n)

	fmt.Println("Введите сами элементы")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	fmt.Printf("Сумма эелементов %d\n", sum)
}

func problem2() {

	fmt.Println("Условие задачи:\n Удалить элемент по заданному индексу из слайса.\n ")

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

func problem3() {

	fmt.Println("Условие задачи:\n Найти максимальное и минимальное значение в слайсе.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	mn, mx := 1000000000, -1000000000

	for i := 0; i < n; i++ {

		if a[i] < mn {
			mn = a[i]
		}

		if a[i] > mx {
			mx = a[i]
		}
	}

	fmt.Printf("Минимум в слайсе: %d\nМаксимум в слайсе: %d\n", mn, mx)
}

func problem4() {

	fmt.Println("Условие задачи:\n Проверить, содержит ли слайс заданный элемент.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println("Какое число проверить?")
	var (
		num int
		ok  bool
	)
	fmt.Scan(&num)

	for i := 0; i < n; i++ {
		if num == a[i] {
			fmt.Printf("Такое число есть, индекс: %d", i+1)
			ok = true
			break
		}
	}
	if !ok {
		fmt.Println("Нет такого числа!")
	}
}

func problem5() {

	fmt.Println("Условие задачи:\n Объединить два слайса в один.\n ")

	var n, m int
	fmt.Println("Введите количество элементов ПЕРВОГО и ВТОРОГО массива")
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, m)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println("Введите элементы второго массива через пробел")
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	c := make([]int, 0, n+m)
	for i := 0; i < n; i++ {
		c = append(c, a[i])
	}
	for i := 0; i < m; i++ {
		c = append(c, b[i])
	}

	fmt.Println("Длина нового слайса", len(c), "и сам слайс", c)
}

func problem6() {

	fmt.Println("Условие задачи:\n Отфильтровать чётные числа из слайса.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	evens := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			evens = append(evens, a[i])
		}
	}
	fmt.Println("Четные числа:", evens)
}

func problem7() {

	fmt.Println("Условие задачи:\n Развернуть слайс в обратном порядке.\n ")

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

func problem8() {

	fmt.Println("Условие задачи:\n Удалить дубликаты из слайса.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	us := make(map[int]int)

	for i := 0; i < len(a); {
		if us[a[i]] > 0 {
			a = append(a[:i], a[i+1:]...)
			continue
		}
		us[a[i]]++
		i++
	}

	fmt.Println("Слайс без дубликатов", a)
}

func problem9() {

	fmt.Println("Условие задачи:\n Удалить все нули из слайса целых чисел.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < len(a); {
		if a[i] == 0 {
			a = append(a[:i], a[i+1:]...)
			continue
		}
		i++
	}
	fmt.Println("Массив без нулей:", a)
}

func problem10() {

	fmt.Println("Условие задачи:\n Увеличить каждый элемент слайса на 1.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		a[i]++
	}
	fmt.Println("Новый слайс", a)
}

func problem11() {

	fmt.Println("Условие задачи:\n Найти среднее значение всех элементов слайса.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sum := 0

	for i := 0; i < n; i++ {
		sum += a[i]
	}
	fmt.Printf("Среднее значение: %.2f", float64(sum)/float64(n))
}

func problem12() {

	// 12
	// 10
	// 54  7 12 0 1 4  -7 -9 12 4
	fmt.Println("Условие задачи:\n Отсортировать слайс по возрастанию (вручную, без sort).\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	mergeSort(a, 0, len(a)-1)

	fmt.Println("Отсортированный слайс:", a)
}

func problem13() {

	fmt.Println("Условие задачи:\n Сравнить два слайса на равенство.\n ")

	var n, m int
	fmt.Println("Введите количество элементов ПЕРВОГО и ВТОРОГО массива")
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, m)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println("Введите элементы второго массива через пробел")
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	if n != m {
		fmt.Println("Слайсы не равны!")
		return
	}
	mergeSort(a, 0, len(a)-1)
	mergeSort(b, 0, len(b)-1)

	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			fmt.Println("Слайсы не равны!")
			return
		}
	}
	fmt.Println("Слайсы одинаковы!")
}

func problem14() {

	fmt.Println("Условие задачи:\n Сдвинуть все элементы слайса на один влево.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 1; i < n; i++ {
		swap(&a[i], &a[i-1])
	}
	fmt.Println("Сдвиг влево на 1:", a)
}

func problem15() {

	fmt.Println("Условие задачи:\n Разбить слайс на чанки фиксированного размера.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Введите элементы первого массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println("Введите размер чанка:")
	var chSize int
	fmt.Scan(&chSize)
	chunks := make([][]int, 0, n/chSize+1)

	for i := 0; i < n; i++ {
		if i%chSize == 0 {
			chunks = append(chunks, make([]int, 0, chSize))
		}
		chunks[len(chunks)-1] = append(chunks[len(chunks)-1], a[i])
	}

	fmt.Println("Разбиение:")
	for i := 0; i < len(chunks); i++ {
		fmt.Printf("Чанк %d\n", i+1)
		for j := 0; j < len(chunks[i]); j++ {
			fmt.Print(chunks[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {

	for {
		fmt.Println("\nВведите номер задачи которую хотите запустить?")
		fmt.Println("ПРЕДУПРЕЖДЕНИЕ! - ВЫХОД ТОЛЬКО С КОМБИНАЦИЕЙ ctrl + c")

		var problem int
		fmt.Scan(&problem)

		switch problem {
		case 1:
			problem1()
		case 2:
			problem2()
		case 3:
			problem3()
		case 4:
			problem4()
		case 5:
			problem5()
		case 6:
			problem6()
		case 7:
			problem7()
		case 8:
			problem8()
		case 9:
			problem9()
		case 10:
			problem10()
		case 11:
			problem11()
		case 12:
			problem12()
		case 13:
			problem13()
		case 14:
			problem14()
		case 15:
			problem15()
		default:
			fmt.Println("Такого номера не существует, повторите заново!")
		}
	}
}
