package main

import (
	"facade/src"
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Фасад - структурный паттерн проектирования, позволяющий скрыть сложность системы путём сведения всех возможных внешних
	вызовов к одному объекту, делегирующему их соответствующим объектам системы.
*/

/*
Плюсы: упрощает использование сложной подсистемой путем предоставления простого интерфейса.
Минусы: рискует сать божественным классом, что усложнит поддержку кода.
*/

/*
В данном случае мы применем паттерн фасад на примере работы интернет магазина. Сведем весь функционал к вызову одной функци -
Facade(). Наш магазин будет проверять баланс клиента на предмет возможности покупки товара; город клиеента,
на предмет возможности доставки туда, и сам товар на пермдмет его наличия на складе. В случае успеха - товар будет продан,
со счета клиента спишется нужная сумма.
*/

func main() {
	facade := src.NewFacade("Kazan", "Dmitriy", 130000)

	err := facade.Sell("Samsung")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = facade.Sell("Dyson")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = facade.Sell("Samsung")

	if err != nil {
		fmt.Println(err)
		return
	}
}
