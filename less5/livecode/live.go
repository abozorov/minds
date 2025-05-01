package main

import (
	"errors"
	"fmt"
)

//func sayHello() {
//	fmt.Println("Hello, Go!")
//
//	sayHi()
//}
//
//func sayHi() {
//	fmt.Println("Hi!")
//}

//func IncrementValue(num int) {
//	if num >= 10 {
//		fmt.Println("The number is 10")
//		return
//	}
//
//	IncrementValue(num + 1)
//}

//func counter() func() int {
//	count := 0
//	return func() int {
//		count++
//		return count
//	}
//}

//func add(a, b, c int, name string) {
//	fmt.Println(a + b + c)
//	fmt.Println(name)
//}

//func sum(numbers ...int) {
//	total := 0
//	for _, num := range numbers {
//		total += num
//	}
//
//	fmt.Printf("Total sum: %d\n", total)
//	//return total
//}

func double(n int) int {
	return n * 2
}

func doubleV2(a int, b float32) (int, float32) {
	return a * 2, b * 2
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func minMax(nums []int) (min int, max int) {
	min, max = nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return
}

func greet(name string, age int) {
	fmt.Printf("Привет, %s! Тебе %d лет\n", name, age)
}

func main() {

	//nums := []int{12, 234, 235, 123, 45, 52, 3}
	//
	//mn, mx := minMax(nums)
	//fmt.Println(mn, mx)

	//num := double(5)
	//fmt.Println(num)
	//
	//num1, num2 := doubleV2(4, 5.6)
	//fmt.Println(num1, num2)

	//result, err := divide(5, 0)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(result)

	//name := "Vasya"
	//age := 10
	//
	//fmt.Printf("Name: %s. Age = %d\n", name, age)
	//info := fmt.Sprintf("Name: %s. Age = %d\n", name, age)
	//fmt.Println(info)

	//sum()
	//sum(1)
	//sum(14, 223)
	//sum(1, 2, 23, 33, 44, 55, 66)
	//
	//var tmp any
	//tmp = 1
	//fmt.Println(tmp)
	//tmp = "Vasya"
	//fmt.Println(tmp)
	//tmp = true
	//fmt.Println(tmp)

	//name := "Vasya"
	//age := 12
	//greet(name, age)
	//add(1, 2, 3, name)

	//c := counter()
	//fmt.Println(c())
	//fmt.Println(c())
	//fmt.Println(c())

	// FILO - first In Last out
	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//defer fmt.Println("3")
	//defer fmt.Println("4")
	//defer fmt.Println("5")

	//add := func(a, b int) int {
	//	return a + b
	//}
	//fmt.Println(add(3, 4))

}
