package src

// конкретный заказ в город Москва
type OrderToMoscow struct {
	address string
	product string
	amount  int
	price   int
}

// устанавливает адрес доставки заказа
func (o *OrderToMoscow) SetAddress() {
	o.address = "Moscow, Tverskaya st."

}

// устанавливает товар в заказе
func (o *OrderToMoscow) SetProduct() {
	o.product = "Apple Macbook pro 13.3"

}

// устанавливает количество товара
func (o *OrderToMoscow) SetAmount() {
	o.amount = 1

}

// устанавливает цену заказа
func (o *OrderToMoscow) SetPrice() {
	o.price = 133800
}

// возвращает созданный заказ
func (o *OrderToMoscow) GetOrder() Order {
	return Order{
		address: o.address,
		product: o.product,
		amount:  o.amount,
		price:   o.price,
	}
}
