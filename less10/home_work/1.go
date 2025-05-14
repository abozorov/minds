package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	Products []*Product
}

// Добавляет товар в инвентарь
func (i *Inventory) AddProduct(p *Product) {
	i.Products = append(i.Products, p)
}

// Находит товар по ID
func (i Inventory) GetProductByID(id int) (*Product, error) {
	for _, v := range i.Products {
		if v.ID == id {
			return v, nil
		}
	}
	return &Product{}, errors.New(fmt.Sprintf("Product with %d Not found", id))
}

// Обновляет количество товара
func (i *Inventory) UpdateQuantity(id, newQuantity int) error {
	for _, v := range i.Products {
		if v.ID == id {
			v.Quantity = newQuantity
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Product with %d Not found", id))
}

// Считает общую стоимость всех товаров
func (i Inventory) GetTotalValue() float64 {
	sum := 0.0

	for _, v := range i.Products {
		sum += v.Price * float64(v.Quantity)
	}
	return sum
}

// Возвращает товары с количеством ниже порога
func (i Inventory) GetLowStockProducts(threshold int) []*Product {
	stockProducts := make([]*Product, 0, len(i.Products))

	for _, v := range i.Products {
		if v.Quantity < threshold {
			stockProducts = append(stockProducts, v)
		}
	}
	return stockProducts
}

// Вывод инвентаря
func (i Inventory) Print() {
	fmt.Println("\n________________________________")
	for _, p := range i.Products {
		fmt.Printf("ID: %d\nНазвание: %s\nЦена: %.2f ₽\nКоличество: %d\n\n", p.ID, p.Name, p.Price, p.Quantity)
	}
}

func main() {

	//	Создайте систему учета товаров в магазине:
	//
	// 1. Создайте структуру Product с полями ID, Name, Price, Quantity
	// 2. Создайте структуру Inventory с полем Products (слайс указателей на Product)
	// 3. Реализуйте методы для Inventory:
	//   - AddProduct(*Product) - добавляет товар в инвентарь
	//   - GetProductByID(id int) *Product - находит товар по ID
	//   - UpdateQuantity(id, newQuantity int) bool - обновляет количество товара
	//   - GetTotalValue() float64 - считает общую стоимость всех товаров
	//   - GetLowStockProducts(threshold int) []*Product - возвращает товары с количеством ниже порога
	// 4. Проверьте работу методов на разных сценариях в main

	products := Inventory{
		[]*Product{
			{ID: 1, Name: "Ноутбук Lenovo IdeaPad 3", Price: 74999.99, Quantity: 12},
			{ID: 2, Name: "Смартфон Samsung Galaxy A52", Price: 54999.50, Quantity: 30},
			{ID: 3, Name: "Монитор LG UltraGear 27\"", Price: 22999.00, Quantity: 8},
			{ID: 4, Name: "Мышка Logitech M185", Price: 1599.90, Quantity: 50},
			{ID: 5, Name: "Клавиатура Redragon Kumara", Price: 3199.00, Quantity: 25},
			// {ID: 6, Name: "Жесткий диск Seagate 1TB", Price: 4299.99, Quantity: 15},
			// {ID: 7, Name: "Наушники JBL Tune 500", Price: 2799.00, Quantity: 20},
			// {ID: 8, Name: "Внешний SSD Samsung T7 500GB", Price: 9999.00, Quantity: 10},
			// {ID: 9, Name: "Веб-камера Logitech C920", Price: 7999.00, Quantity: 14},
			// {ID: 10, Name: "Игровое кресло DXRacer", Price: 23999.99, Quantity: 5},
			// {ID: 11, Name: "Планшет Xiaomi Pad 5", Price: 37999.00, Quantity: 9},
			// {ID: 12, Name: "Смарт-часы Apple Watch SE", Price: 25999.99, Quantity: 18},
			// {ID: 13, Name: "Колонка Marshall Emberton", Price: 11999.00, Quantity: 7},
			// {ID: 14, Name: "Маршрутизатор TP-Link Archer AX10", Price: 5999.00, Quantity: 16},
			// {ID: 15, Name: "USB-хаб Orico 4 порта", Price: 1199.00, Quantity: 40},
		},
	}
	// Вывод информации о каждом товаре
	products.Print()

	products.AddProduct(&Product{ID: 6, Name: "Жесткий диск Seagate 1TB", Price: 4299.99, Quantity: 15})
	products.AddProduct(&Product{ID: 7, Name: "Наушники JBL Tune 500", Price: 2799.00, Quantity: 20})
	products.AddProduct(&Product{ID: 8, Name: "Внешний SSD Samsung T7 500GB", Price: 9999.00, Quantity: 10})

	// Вывод информации о каждом товаре
	products.Print()

	fmt.Println(products.GetProductByID(155))
	p, err := products.GetProductByID(1)
	fmt.Println(*p, err)
	p, err = products.GetProductByID(7)
	fmt.Println(*p, err)

	// Вывод информации о каждом товаре
	products.Print()

	fmt.Println(products.UpdateQuantity(15, 25))
	fmt.Println(products.UpdateQuantity(8, 1))
	fmt.Println(products.UpdateQuantity(1, 30))

	// Вывод информации о каждом товаре
	products.Print()

	fmt.Printf("SUMM: %.2f", products.GetTotalValue())

	x := Inventory{products.GetLowStockProducts(25)}
	x.Print()
}
