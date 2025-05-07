package main

import "fmt"

type UserName string

func main() {
	var name UserName
	fmt.Println("Введите ИМЯ: ")

	fmt.Scan(&name)

	fmt.Printf("ИМЯ %s", name)
}
