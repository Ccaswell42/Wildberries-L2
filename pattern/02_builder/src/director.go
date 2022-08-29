package src

//Создание заказа контролируется объектом-распорядителем (Director), которому нужно знать только тип создаваемого объекта.

type Director struct {
	Collector Collector
}

// конструктор

func NewDirector(collector Collector) *Director {
	return &Director{Collector: collector}
}

// устанавливаем соответствующую комплектацию заказа

func (d *Director) SetCollector(collector Collector) {
	d.Collector = collector
}

// создаем сам заказ

func (d *Director) CreateOrder() Order {
	d.Collector.SetAddress()
	d.Collector.SetPrice()
	d.Collector.SetAmount()
	d.Collector.SetProduct()
	return d.Collector.GetOrder()
}
