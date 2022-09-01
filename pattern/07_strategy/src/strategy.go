package src

// Strategy предоставляет интерфейс для алгоритмов сортировки.
type Strategy interface {
	Price(distance, productWeight int)
}
