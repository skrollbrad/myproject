// структура order - будет содержать слайс позиций, общую стоимость заказа и пользователя, id ++
// структура position - название товара/id, количество, стоимость ++
// product struct - id, name, стоимость ++
// создание пользователя - сохраняем в мапу ++
// добавление продукта ++
// функционал корзины(смотрим по id пользователя)
// endpoint получения корзины пользователя по id
// кнопка оформить заказ - всё складываем из корзины, считаем фулл стоимость и после оформления удалить всё из корзины.
package models

import "fmt"

type Order struct {
	Pos            string
	FullPriceOrder float64
	Id             int
}
type Position struct {
	ProductName string
	Count       int
	Price       float64
}
type User struct {
	ID       string
	Username string
	Email    string
}

type Product struct {
	Name  string
	Price float64
}

type Store struct {
	Products map[string]Product
}

func (s *Store) AddProduct(name string, price float64) {
	product := Product{Name: name, Price: price}
	s.Products[name] = product
}

func main() {
	store := Store{
		Products: make(map[string]Product),
	}

	store.AddProduct("Apple", 2.5)
	store.AddProduct("Banana", 1.2)

	fmt.Println(store.Products)
}

// func PersonAdd
