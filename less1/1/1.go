package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Введите ваше имя и возраст:")
	fmt.Scanln(&name, &age)
	fmt.Printf("Привет, %s! через 10 лет тебе будет %d лет.\n", name, age+10)
}
