package main

import (
	"fmt"
	"strconv"
)

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
	fmt.Println("Условие задачи:\n Подсчитать количество вхождений каждого элемента в слайсе и сохранить в map.\n ")

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

	for key, val := range us {
		fmt.Printf("Элемент \"%s\" и кол-во вхождений %d\n", key, val)
	}
}

func problem2() {

	fmt.Println("Условие задачи:\n Инвертировать map[string]int в map[int]string.\n ")

	fmt.Println("Введите жлементы \"Ключ\" \"Значение\" через пробел.")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value int
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}
	reverseMp := make(map[int]string)

	for k, v := range mp {
		reverseMp[v] = k
	}

	fmt.Println("\nПеревернутый мап:\nКлюч - Значение\n ")
	for k, v := range reverseMp {
		fmt.Printf("%d - %s\n", k, v)
	}
}

func problem3() {
	/*
		adasd 123123
		adakf 343
		jdfls 345
		21324 2132222222
		dhjkdf 44
		exit
	*/
	fmt.Println("Условие задачи:\n Проверить, существует ли ключ в map.\n ")

	fmt.Println("Введите элементы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - целое число")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value int
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}
	fmt.Println("На каждой строке вводите ключ который хотите проверить.")
	fmt.Println("Для выхода введите команду exit")
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			value, exist := mp[key]

			if exist {
				fmt.Printf("Такой ключ есть, вот значение %d\n", value)
			} else {
				fmt.Println("Такого ключа нет")
			}
		} else {
			break
		}
	}

}

func problem4() {

	fmt.Println("Условие задачи:\n Удалить элемент по ключу из map.\n ")

	fmt.Println("Введите элементы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - целое число")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value int
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}
	fmt.Println("Мапа до удаления", mp, "\nВведите ключ удаляемого элемента: ")
	var deleted string
	fmt.Scan(&deleted)
	delete(mp, deleted)
	fmt.Println("Мапа после удаления", mp)
}

func problem5() {

	fmt.Println("Условие задачи:\n Объединить два map[string]int, суммируя значения одинаковых ключей.\n ")

	fmt.Println("Введите элементы первой мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - целое число")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value int
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}

	fmt.Println("Введите элементы второй мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - целое число")
	fmt.Println("После ввода всех данных введите команду exit")

	mp2 := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value int
			fmt.Scan(&value)
			mp2[key] = value
		} else {
			break
		}
	}

	for k, v := range mp2 {
		mp[k] += v
	}

	fmt.Println("Объединенный мап:", mp)
}

func problem6() {

	fmt.Println("Условие задачи:\n Найти ключ с максимальным значением.\n ")

	fmt.Println("Введите элементы первой мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - целое число")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value int
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}
	var mxKey string

	for k, v := range mp {
		if mxKey == "" {
			mxKey = k
		} else if mp[mxKey] < v {
			mxKey = k
		}
	}

	fmt.Printf("Максимальный ключ %s и его значение %d\n", mxKey, mp[mxKey])
}

func problem7() {

	fmt.Println("Условие задачи:\n Найти все ключи с определённым значением.\n ")

	fmt.Println("Введите элементы первой мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - строка")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]string)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value string
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}
	fmt.Println("Введите значение ключи которой хотите найти:")
	var val string
	fmt.Scan(&val)
	a := make([]string, 0, len(mp))

	for k, v := range mp {

		if v == val {
			a = append(a, k)
		}
	}

	fmt.Println("Ключи с таким значением", a)
}

func problem8() {

	fmt.Println("Условие задачи:\n Подсчитать количество уникальных значений в map.\n ")

	fmt.Println("Введите элементы первой мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - строка")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]string)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value string
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}
	us := make(map[string]int)

	for _, v := range mp {
		us[v]++
	}
	ans := 0
	for _, v := range mp {

		if us[v] == 1 {
			ans++
		}
	}

	fmt.Println("Кол-во уникальных значений:", ans)
}

func problem9() {

	fmt.Println("Условие задачи:\n Скопировать map в новую переменную.\n ")

	fmt.Println("Введите элементы первой мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - строка")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]string)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value string
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}
	cmp := make(map[string]string)

	for k, v := range mp {
		cmp[k] = v
	}

	fmt.Println("Копия мапы", cmp)
}

func problem10() {

	fmt.Println("Условие задачи:\n Удалить все элементы map.\n ")

	fmt.Println("Введите элементы первой мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - строка")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]string)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value string
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}

	fmt.Println("Сам мап", mp)

	for k := range mp {
		delete(mp, k)
	}

	fmt.Println("Пустой мап", mp)
}

func problem11() {

	fmt.Println("Условие задачи:\n Создать map, где ключ — слово, а значение — его длина.\n ")

	fmt.Println("Введите строки. Каждую строку на новой строке.")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)

		if key != "exit" {
			mp[key] = len(key)
		} else {
			break
		}
	}

	for k, v := range mp {
		fmt.Printf("Строка \"%s\" его длина %d\n", k, v)
	}
}

func problem12() {

	fmt.Println("Условие задачи:\n Построить map, где ключ — первая буква слова, а значение — количество таких слов.\n ")

	var n int
	fmt.Println("Введите количество элементов массива")
	fmt.Scan(&n)
	us := make(map[byte]int)

	fmt.Println("Введите элементы(строки) первого массива через пробел")
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		us[s[0]]++
	}

	for k, v := range us {
		fmt.Printf("Первая буква %s и кол-во слов %d\n", string(k), v)
	}
}

func problem13() {

	fmt.Println("Условие задачи:\n Преобразовать map[int]string в слайс структур.\n ")

	fmt.Println("Введите элементы первой мапы \"Ключ\" \"Значение\" через пробел. Ключ - число, значение - строка")
	fmt.Println("После ввода всех данных введите команду exit")

	mp := make(map[int]string)
	n := 0

	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value string
			fmt.Scan(&value)
			k, _ := strconv.Atoi(key)
			mp[k] = value

			if n <= k {
				n = k + 1
			}
		} else {
			break
		}
	}
	a := make([]string, n)

	for k, v := range mp {
		a[k] = v
	}

	fmt.Println("Массив:")
	for i := 0; i < n; i++ {
		fmt.Printf("Индекс в массиве %d и строка \"%s\"\n", i, a[i])
	}
	/*
		1 anush
		2 hhh
		7 jasgjha
		4 gahgh
		3 aaaa
		exit
	*/
}

func problem14() {

	fmt.Println("Условие задачи:\n Найти пересечение двух map по ключам.\n ")

	fmt.Println("Введите элементы первой мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - целое число")
	fmt.Println("После ввода всех данных введите команду exit")
	mp := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value int
			fmt.Scan(&value)
			mp[key] = value
		} else {
			break
		}
	}

	fmt.Println("Введите элементы второй мапы \"Ключ\" \"Значение\" через пробел. Ключ - строк, значение - целое число")
	fmt.Println("После ввода всех данных введите команду exit")
	mp2 := make(map[string]int)
	for {
		var key string
		fmt.Scan(&key)
		if key != "exit" {
			var value int
			fmt.Scan(&value)
			mp2[key] = value
		} else {
			break
		}
	}
	fmt.Println("ЗАДАЧА НЕ ПОНЯТНА!")
	// for k1 := range mp {
	// 	for k2 := range mp2 {
	// 		if k1 == k2 {

	// 		}
	// 	}
	// }

}

func problem15() {

	fmt.Println("Условие задачи:\n Отсортировать ключи map и вывести пары по возрастанию ключей.\n ")

	fmt.Println("ЗАДАЧА НЕ ПОНЯТНА, ТУТ КЛЮЧИ ЭТО ЦЕЛЫЕ ЧИСЛА?!")
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
