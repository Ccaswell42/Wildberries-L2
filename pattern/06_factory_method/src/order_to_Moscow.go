package src

import "fmt"

// OrderToMoscow креализует интерфейс Collector
type OrderToMoscow struct {
	address string
	product string
	amount  int
	price   int
}

// NewOrderToMoscow конструктор
func NewOrderToMoscow() Collector {
	return OrderToMoscow{
		address: "Moscow, Tverskaya st.",
		product: "Apple Macbook pro 13.3",
		amount:  1,
		price:   133800,
	}

}

// Print выводит информацию об объекте в консоль
func (o OrderToMoscow) Print() {
	fmt.Printf("адрес: %s, товар: %s, количество: %d, цена: %d\n", o.address, o.product, o.amount, o.price)
}
