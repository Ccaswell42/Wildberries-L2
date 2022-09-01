package src

import "fmt"

// Taxi - алгоритм расчета стоимости доставки товара такси
type Taxi struct {
}

func (t Taxi) Price(distance, productWeight int) {
	if productWeight > 100 { // больше 100 кг такси увезти не может
		fmt.Print("доставка с помощью такси невозможна")
		return
	}
	costOfTaxi := 0 // стоимость работу такси для компании
	// расчет стоимости доставки с учетом коэффициентов для алгоритма "такси"
	totalPrice := distance*5 + costOfTaxi + productWeight*7
	fmt.Printf("доставка с помощью такси будет стоить: %d\n", totalPrice)
}
