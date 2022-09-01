package src

import "fmt"

// Courier - алгоритм расчета стоимости доставки товара курьером
type Courier struct {
}

func (c Courier) Price(distance, productWeight int) {
	if productWeight > 25 { // больше 25 кг курьер нести не может
		fmt.Print("доставка курьером невозможна")
		return
	}
	costOfCourier := 200 // стоимость работу курьера для компании
	// расчет стоимости доставки с учетом коэффициентов для алгоритма "курьер"
	totalPrice := distance*2 + costOfCourier + productWeight*5
	fmt.Printf("доставка курьером будет стоить: %d\n", totalPrice)
}
