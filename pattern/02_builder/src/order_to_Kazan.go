package src

// конкретный заказ в город Казань
type OrderToKazan struct {
	address string
	product string
	amount  int
	price   int
}

// устанавливает адрес доставки заказа
func (o *OrderToKazan) SetAddress() {
	o.address = "Kazan, Bauman st."

}

// устанавливает товар в заказе
func (o *OrderToKazan) SetProduct() {
	o.product = "Кроссовки Reebok Rewind Run"

}

// устанавливает количество товара
func (o *OrderToKazan) SetAmount() {
	o.amount = 2

}

// устанавливает цену заказа
func (o *OrderToKazan) SetPrice() {
	o.price = 4750
}

// возвращает созданный заказ
func (o *OrderToKazan) GetOrder() Order {
	return Order{
		address: o.address,
		product: o.product,
		amount:  o.amount,
		price:   o.price,
	}
}
