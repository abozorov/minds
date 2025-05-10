package main

import (
	"fmt"
	"time"
)

type ActionI interface {
	Do()
}

func LogAction(a ActionI) {
	a.Do()
}

// Что то типо таймера
type Timer struct {
	StartTime time.Time
}

func (t *Timer) Do() {
	if t.StartTime.Second() == 0 {
		t.StartTime = time.Now()
		fmt.Println("Засекаем время!")
		return
	}

	fmt.Printf("Прошло времени %s\n", time.Now().Sub(t.StartTime))
	t.StartTime = time.Time{}
}

// Структура для бибикания машины
type Car struct{}

func (c Car) Do() {
	fmt.Println("Машина бибикает!")
}

func mimiTimer(a ActionI) {
	fmt.Println("Программа засекатель времени!\n - Для замуска введите команду do\n - Для выхода введите команду exit !")
	for {
		var cmd string
		ok := false
		fmt.Scan(&cmd)
		switch cmd {
		case "do":
			LogAction(a)
		case "exit":
			ok = true
		default:
			fmt.Println("Неверная команда!")
		}

		if ok {
			break
		}
	}
}

func main() {
	var a ActionI

	// Запускаем функцию для работы таймера
	a = &Timer{}
	mimiTimer(a)
	
	// Побибикаем
	a = Car{}
	LogAction(a)
	LogAction(a)
	LogAction(a)
	LogAction(a)
	LogAction(a)

	// Еще раз запускаем функцию для работы таймера
	a = &Timer{}
	mimiTimer(a)
}