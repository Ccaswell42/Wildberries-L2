package main

import "builder/src"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

func main() {
	// определили заказы, которые возможны
	kazan := src.GetCollector("Kazan")
	moscow := src.GetCollector("Moscow")

	// создали базовый завод с выбранной по умолчанию комплектацией "kazan", и делаем соответствующие заказы
	factory := src.NewFactory(kazan)
	kazanOrder := factory.CreateOrder()
	kazanOrder.Print()

	// изменили формат на другую комплектацию (другой заказ)
	// и делаем соответствуюший заказ и выводим результат на экран
	factory.SetCollector(moscow)
	moscowOrder := factory.CreateOrder()
	moscowOrder.Print()

}
