package src

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (f *Factory) SetCollector(collector Collector) {
	f.Collector = collector
}

func (f *Factory) CreateOrder() Order {
	f.Collector.SetAddress()
	f.Collector.SetPrice()
	f.Collector.SetAmount()
	f.Collector.SetProduct()
	return f.Collector.GetOrder()
}
