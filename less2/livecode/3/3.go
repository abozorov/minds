package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}
