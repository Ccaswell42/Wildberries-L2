package src

type OrderToKazan struct {
	address string
	product string
	amount  int
	price   int
}

func (o *OrderToKazan) SetAddress() {
	o.address = "Bauman st."

}
func (o *OrderToKazan) SetProduct() {
	o.product = "Кроссовки Reebok Rewind Run"

}
func (o *OrderToKazan) SetAmount() {
	o.amount = 2

}
func (o *OrderToKazan) SetPrice() {
	o.price = 4750
}

func (o *OrderToKazan) GetOrder() Order {
	return Order{
		address: o.address,
		product: o.product,
		amount:  o.amount,
		price:   o.price,
	}
}
