package src

import "fmt"

// Truck - алгоритм расчета стоимости доставки товара грузовой машиной
type Truck struct {
}

func (t Truck) Price(distance, productWeight int) {
	costOfTruck := 400 // стоимость работу грузовой машины для компании
	// расчет стоимости доставки с учетом коэффициентов для алгоритма "Truck"
	totalPrice := distance*4 + costOfTruck + productWeight*3
	fmt.Printf("доставка грузовым автомобилем будет стоить: %d\n", totalPrice)
}
