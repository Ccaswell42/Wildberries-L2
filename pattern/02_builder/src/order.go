package src

import "fmt"

var (
	ToMoscowOrderType = "Moscow"
	ToKazanOrderType  = "Kazan"
)

// абстрактный объект "заказ", который мы будем создавать

type Order struct {
	address string
	product string
	amount  int
	price   int
}

// Метод принт для вывода информации о зкакзе в консоль
func (o *Order) Print() {
	fmt.Printf("adress: %s, product: %s, amount:%d, price: %d\n", o.address, o.product, o.amount, o.price)
}

// Интерфейс сборщик заказа
type Collector interface {
	SetAddress()
	SetProduct()
	SetAmount()
	SetPrice()
	GetOrder() Order
}

// возвращает конкретный заказ по названию города

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case ToMoscowOrderType:
		return &OrderToMoscow{}
	case ToKazanOrderType:
		return &OrderToKazan{}

	}

}
