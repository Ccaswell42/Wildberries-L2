package src

import "fmt"

var (
	ToMoscowOrderType = "Moscow"
	ToKazanOrderType  = "Kazan"
)

type Collector interface {
	SetAddress()
	SetProduct()
	SetAmount()
	SetPrice()
	GetOrder() Order
}

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

type Order struct {
	address string
	product string
	amount  int
	price   int
}

func (o *Order) Print() {
	fmt.Printf("adress: %s, product: %s, amount:%d, price: %d\n", o.address, o.product, o.amount, o.price)
}
