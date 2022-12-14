package main

import "builder/src"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн строитель - порождающий шаблон проектирования, позволяющий поэтапно создать сложный составной объект.
Отделяет конструирование сложного объекта от его представления так, что в результате одного и того же процесса
конструирования могут получаться разные представления.
*/

/*
Плюсы :
-позволяет изменять внутреннее представление продукта;
-изолирует код, реализующий конструирование и представление;
-дает более тонкий контроль над процессом конструирования.
*/
/*
Минусы :
-алгоритм создания сложного объекта не должен зависеть от того, из каких частей состоит объект
 и как они стыкуются между собой;
-процесс конструирования должен обеспечивать различные представления конструируемого объекта.
*/

// Реализуем работу паттерна 'строитель на примере создания сложного объекта - заказ.

func main() {
	// определили комплектации заказов, которые возможны
	kazan := src.GetCollector("Kazan")
	moscow := src.GetCollector("Moscow")

	director := src.NewDirector(kazan)   // <- создаем директора со строителем "заказ в Казань"
	kazanOrder := director.CreateOrder() // <- создаем сам заказ
	kazanOrder.Print()                   // <- выводим результат в консоль

	// меняем комплектацию  на "заказ в Москву"
	director.SetCollector(moscow)         // <- меняем комплектацию  на "заказ в Москву"
	moscowOrder := director.CreateOrder() // <- создаем новый заказ
	moscowOrder.Print()                   // <- выводим результат в консоль

}
