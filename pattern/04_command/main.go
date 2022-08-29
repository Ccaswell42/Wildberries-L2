package main

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Паттерн Command относится к поведенческим паттернам уровня объекта.
Паттерн Command позволяет представить запрос в виде объекта.
Из этого следует, что команда - это объект. Такие запросы, например, можно ставить в очередь, отменять или возобновлять.
В этом паттерне мы оперируем следующими понятиями: Command - запрос в виде объекта на выполнение;
Receiver - объект-получатель запроса, который будет обрабатывать нашу команду; Invoker - объект-инициатор запроса.
Паттерн Command отделяет объект, инициирующий операцию, от объекта, который знает, как ее выполнить.
Единственное, что должен знать инициатор, это как отправить команду.
*/

/*
Плюсы :
-Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
-Позволяет собирать сложные команды из простых.

Минусы:
-Усложняет код программы из-за введения множества дополнительных классов.
*/

/*
	В данном примере мы создали структуру ресторана (Receiver) который принимает задания от списка заданий и выполняет
команды по доставке пиццы, доставке суши, и возврату курьеров обратно в ресторан. Сами функции отделены от списка заданий
и инкапсулированы в общий интерфейс.
*/

// Command предоставляет коммандный интерфейс
type Command interface {
	Execute()
}

// Restaurant - структура, имеющая кассу и количество незанятых курьеров. Receiver
type Restaurant struct {
	Couriers int
	Cash     int
}

// NewRestaurant - конструктор
func NewRestaurant() *Restaurant {
	return &Restaurant{
		Couriers: 10,
		Cash:     0,
	}
}

// методы структуры Restaurant, реализующие наши команды

func (r *Restaurant) MakePizzaDelivery() Command {
	return &PizzaCmd{
		restaurant: r,
		price:      700,
	}
}

func (r *Restaurant) MakeSushiDelivery() Command {
	return &SushiCmd{
		restaurant: r,
		price:      1300,
	}
}

func (r *Restaurant) MakeComeBack() Command {
	return &ComeBackCmd{
		restaurant: r,
	}
}

// PizzaCmd включает в себя ссылку на ресторан и цену пиццы. Команда
type PizzaCmd struct {
	restaurant *Restaurant
	price      int
}

// Execute - наша команда для PizzaCmd
func (p *PizzaCmd) Execute() {
	p.restaurant.Couriers -= 1
	p.restaurant.Cash += p.price
	fmt.Printf("заказ [доставка пиццы] оформлен, количество курьеров: %d, касса: %d\n", p.restaurant.Couriers, p.restaurant.Cash)
}

// SushiCmd включает в себя ссылку на ресторан и цену суши. Команда
type SushiCmd struct {
	restaurant *Restaurant
	price      int
}

// Execute - наша команда для SushiCmd
func (s *SushiCmd) Execute() {
	s.restaurant.Couriers -= 1
	s.restaurant.Cash += s.price
	fmt.Printf("заказ [доставка суши] оформлен, количество курьеров: %d, касса: %d\n", s.restaurant.Couriers, s.restaurant.Cash)

}

// ComeBackCmd и метод Execute - возвращает одного курьера обратно в рсторан. Команда
type ComeBackCmd struct {
	restaurant *Restaurant
}

func (c *ComeBackCmd) Execute() {
	c.restaurant.Couriers += 1
	fmt.Printf("команда [курьер вернулся] выполнена, количество курьеров: %d\n", c.restaurant.Couriers)
}

// CommandList - список комманд, Invoker.
type CommandList struct {
	commands []Command
}

// ExecuteCommands исполняет команды из списка по очереди
func (c *CommandList) ExecuteCommands() {
	for _, val := range c.commands {
		val.Execute()
	}

}
func main() {
	r := NewRestaurant() // <- инициализируем наш ресторан
	tasks := CommandList{
		commands: []Command{
			r.MakePizzaDelivery(),
			r.MakePizzaDelivery(),
			r.MakePizzaDelivery(),
			r.MakeSushiDelivery(),
			r.MakeComeBack(),
			r.MakeComeBack(),
		},
	} // <- создаем список команд
	for _, val := range tasks.commands {
		val.Execute()
		time.Sleep(2 * time.Second)
	} // <-  исполняем наши команды
}
