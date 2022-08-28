package src

type OrderToMoscow struct {
	address string
	product string
	amount  int
	price   int
}

func (o *OrderToMoscow) SetAddress() {
	o.address = "Tverskaya st."

}
func (o *OrderToMoscow) SetProduct() {
	o.product = "Apple Macbook pro 13.3"

}
func (o *OrderToMoscow) SetAmount() {
	o.amount = 1

}
func (o *OrderToMoscow) SetPrice() {
	o.price = 133800
}

func (o *OrderToMoscow) GetOrder() Order {
	return Order{
		address: o.address,
		product: o.product,
		amount:  o.amount,
		price:   o.price,
	}
}
