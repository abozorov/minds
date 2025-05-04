package main

import (
	"fmt"
)

// #1 Students
type Student struct {
	Name  string
	Age   int
	Group string
	GPA   float64
}

func newStudent(name string, age int, group string, gpa float64) Student {
	return Student{
		Name:  name,
		Age:   age,
		Group: group,
		GPA:   gpa,
	}
}

func printStudent(s Student) {
	fmt.Printf("Студент: <%s>, Возраст: <%d>, Группа: <%s>, Средний балл: <%.2f>\n", s.Name, s.Age, s.Group, s.GPA)
}

func printExcellentStudents(students []Student) {
	for _, v := range students {
		if v.GPA > 3.5 {
			printStudent(v)
		}
	}
}

// #3 City
type City struct {
	Name       string
	Population int
	IsCapital  bool
}

func newCity(name string, population int, isCapital bool) City {
	return City{
		Name:       name,
		Population: population,
		IsCapital:  isCapital,
	}
}

// #4 Game
type Game struct {
	Title       string
	Platform    string
	HoursPlayed int
}

func newGame(title, platform string, hoursPlayed int) Game {
	return Game{
		Title:       title,
		Platform:    platform,
		HoursPlayed: hoursPlayed,
	}
}

func totalPlaytime(games []Game) int {
	total := 0
	for _, v := range games {
		total += v.HoursPlayed
	}
	return total
}

// #5 Order
type Order struct {
	Item         string
	Quantity     int
	PricePerUnit float64
}

func newOrder(item string, quantity int, pricePerUnit float64) Order {
	return Order{
		Item: item,
		Quantity: quantity,
		PricePerUnit: pricePerUnit,
	}
}

func main() {

	// 	1. Создайте структуру Student
	//  Поля: Name (string), Age (int), Group (string), GPA (float64).
	//  - Создайте слайс из 3 студентов.
	//  - Выведите всех студентов в формате:
	// Name: <...>, Age: <...>, Group: <...>, GPA: <...>
	students := make([]Student, 0, 100)
	students = append(students,
		newStudent("Bob", 18, "B", 3),
		newStudent("Odil", 19, "A", 5),
		newStudent("Kate", 20, "A", 4),
		newStudent("Алексей Смирнов", 20, "БИ-201", 4.5),
		newStudent("Мария Иванова", 19, "ФИ-102", 3.2),
		newStudent("Дмитрий Кузнецов", 21, "ИТ-301", 4.7),
		newStudent("Екатерина Орлова", 22, "ЭК-401", 2.9),
		newStudent("Иван Петров", 20, "БИ-201", 3.8),
		newStudent("Светлана Васильева", 23, "ЮР-101", 3.3),
		newStudent("Никита Соколов", 19, "ИТ-101", 4.9),
	)
	fmt.Println("Все студенты:")
	for _, v := range students {
		printStudent(v)
	}

	// 2. Фильтрация по оценке
	//  - Используйте структуру Student из задачи 1.
	//  - Напишите функцию printExcellentStudents(students []Student), которая
	// выводит только студентов с GPA выше 3.5.
	fmt.Println("Лучшие из них:")
	printExcellentStudents(students)

	// 3. Создайте структуру City
	//  Поля: Name (string), Population (int), IsCapital (bool).
	//  - Создайте слайс из 4 городов.
	//  - Посчитайте и выведите количество столиц (IsCapital == true) и общее
	// население всех городов.
	cities := make([]City, 0, 100)
	cities = append(cities,
		newCity("Новосибирск", 1620000, false),
		newCity("Берлин", 3850000, true),
		newCity("Чикаго", 2710000, false),
		newCity("Оттава", 1000000, true),
		newCity("Саппоро", 1960000, false),
		newCity("Лиссабон", 505000, true),
		newCity("Казань", 1250000, false),
	)
	capitals := 0
	allPopulation := 0
	for _, v := range cities {
		if v.IsCapital {
			capitals++
		}
		allPopulation += v.Population
	}
	fmt.Printf("\n\nCapitals: %d & All Population: %d", capitals, allPopulation)

	// 4. Создайте структуру Game
	//  Поля: Title (string), Platform (string), HoursPlayed (int).
	//  - Создайте слайс игр.
	//  - Напишите функцию totalPlaytime(games []Game), которая возвращает общее
	// количество часов.
	games := make([]Game, 0, 100)
	games = append(games,
		newGame("The Witcher 3", "PC", 120),
		newGame("God of War", "PlayStation", 45),
		newGame("Stardew Valley", "Switch", 60),
		newGame("Cyberpunk 2077", "PC", 30),
		newGame("The Legend of Zelda: Breath of the Wild", "Switch", 150),
		newGame("Halo Infinite", "Xbox", 25),
		newGame("Hades", "PC", 12),
	)
	fmt.Printf("\n\nTotal wasted time: %d", totalPlaytime(games))

	// 5. Создайте структуру Order
	//  Поля: Item (string), Quantity (int), PricePerUnit (float64).
	//  - Создайте массив заказов.
	//  - Посчитайте и выведите итоговую сумму всех заказов:
	//  Quantity * PricePerUnit для каждого заказа и суммарно.
	orders := make([]Order, 0, 100)
	orders = append(orders,
		newOrder("Ноутбук", 2, 75000.00),
		newOrder("Смартфон", 5, 32000.50),
		newOrder("Наушники", 10, 3500.00),
		newOrder("Монитор", 3, 18999.99),
		newOrder("Клавиатура", 4, 2700.00),
		newOrder("Мышь", 6, 1500.50),
		newOrder("Веб-камера", 2, 5400.00),
		newOrder("Принтер", 1, 12000.00),
		newOrder("USB-хаб", 8, 950.00),
		newOrder("Внешний жёсткий диск", 3, 8900.00),
		newOrder("Микрофон", 2, 6200.75),
		newOrder("Кресло офисное", 2, 13500.00),
		newOrder("Коврик для мыши", 5, 450.00),
		newOrder("Сетевой фильтр", 4, 1100.00),
		newOrder("Колонки", 3, 4300.00),
	)
	sumOfOrders := 0.0
	for _, v := range orders {
		sumOfOrders += float64(v.Quantity) * v.PricePerUnit
	}
	fmt.Printf("\n\nSum of all orders: %.2f", sumOfOrders)
}
