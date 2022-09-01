package src

import "fmt"

// OrderToKazan реализует интерфейс Collector
type OrderToKazan struct {
	address string
	product string
	amount  int
	price   int
}

// NewOrderToKazan - конструктор
func NewOrderToKazan() Collector {
	return OrderToKazan{
		address: "Kazan, Bauman st.",
		product: "Кроссовки Reebok Rewind Run",
		amount:  2,
		price:   4750,
	}

}

// Print выводит информацию об объекте в консоль
func (o OrderToKazan) Print() {
	fmt.Printf("адрес: %s, товар: %s, количество: %d, цена: %d\n", o.address, o.product, o.amount, o.price)
}
