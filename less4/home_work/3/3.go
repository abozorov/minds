package main

import "fmt"

func main() {
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

/*
adasd 123123
adakf 343
jdfls 345
21324 2132222222
dhjkdf 44
adsass 343
sfsuuu 44
suuuuuuuuuu 44
exit
*/