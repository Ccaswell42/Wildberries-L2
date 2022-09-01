package src

import "fmt"

var (
	ToMoscowOrderType = "Moscow"
	ToKazanOrderType  = "Kazan"
)

// Collector предоставляет фабричный интерфейс
type Collector interface {
	Print()
}

// GetCollector - фабричный метод
func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		fmt.Printf("%s такого типа не существует\n", collectorType)
		return nil
	case ToMoscowOrderType:
		return NewOrderToMoscow()
	case ToKazanOrderType:
		return NewOrderToKazan()
	}
}
