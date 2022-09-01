package main

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Паттерн "цепочка вызовов" представляет собой последовательность объектов-обработчиков запроса, которые решают:
обработать запрос или передать его дальше по цепочке.
*/

/*
Плюсы:
-позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса.
-имеется возможность добавления новых объектов-обработчиков в существуюшую цепь.
*/

/*
Минусы: Есть вероятность что запрос не будет обработан ни одним обработчиком и потеряется.
*/

/*
Реализуем паттерн "цепочка вызовов" на примере приготовления пиццы.
*/

// Handler - интерфейс для объектов-обработчиков
type Handler interface {
	Add(pizza *Pizza)
	SetNext(h Handler)
}

// Pizza - объект(запрос), который будем обрабатывать
type Pizza struct {
	Sauce  string
	Meat   string
	Cheese string
}

// Sauce - объект-обработчик, имеет название и указатель на следующий элемент в цепочке.
type Sauce struct {
	Name string
	Next Handler
}

// Add - добавляет соус если его нет, и отправляет нашу пиццу дальше
func (s *Sauce) Add(pizza *Pizza) {
	if pizza.Sauce == "" {
		pizza.Sauce = s.Name
		fmt.Printf("add %s sauce\n", s.Name)
	} else {
		fmt.Println("sauce already added")
	}
	time.Sleep(1 * time.Second)
	if s.Next != nil {
		s.Next.Add(pizza)
	}
}

// SetNext указывает на следующий объект-обработчик в цепочке
func (s *Sauce) SetNext(h Handler) {
	s.Next = h
}

// Meat - объект-обработчик, имеет название и указатель на следующий элемент в цепочке.
type Meat struct {
	Name string
	Next Handler
}

// Add - добавляет мясо если его нет, и отправляет нашу пиццу дальше
func (m *Meat) Add(pizza *Pizza) {
	if pizza.Meat == "" {
		pizza.Meat = m.Name
		fmt.Printf("add %s\n", m.Name)
	} else {
		fmt.Println("meat already added")
	}
	time.Sleep(1 * time.Second)
	if m.Next != nil {
		m.Next.Add(pizza)
	}
}

// SetNext указывает на следующий объект-обработчик в цепочке
func (m *Meat) SetNext(h Handler) {
	m.Next = h
}

// Cheese - объект-обработчик, имеет название и указатель на следующий элемент в цепочке.
type Cheese struct {
	Name string
	Next Handler
}

// Add - добавляет сыр если его нет, и отправляет нашу пиццу дальше
func (c *Cheese) Add(pizza *Pizza) {
	if pizza.Cheese == "" {
		pizza.Cheese = c.Name
		fmt.Printf("add %s\n", c.Name)
	} else {
		fmt.Println("cheese already added")
	}
	time.Sleep(1 * time.Second)
	if c.Next != nil {
		c.Next.Add(pizza)
	}
}

// SetNext указывает на следующий объект-обработчик в цепочке
func (c *Cheese) SetNext(h Handler) {
	c.Next = h
}

func main() {
	sauce := &Sauce{Name: "Tomato"}       // <- устанавливаем объект обработчик с добавлением соуса
	meat := &Meat{Name: "Pepperoni"}      // <- устанавливаем объект обработчик с добавлением мяса
	cheese := &Cheese{Name: "Mozzarella"} // <- устанавливаем объект обработчик с добавлением сыра
	sauce.SetNext(meat)                   // <- устанавливаем следующий объект в цепочке обрабортки запроса
	meat.SetNext(cheese)                  // <- устанавливаем следующий объект в цепочке обрабортки запроса
	pizza := &Pizza{Sauce: "BBQ"}         // <- инициализируем наш запрос
	sauce.Add(pizza)                      // <- запускаем нашу цепоку вызовов

	// выводим полученный результат в консоль
	fmt.Println("ваша пицца готова, ингридиенты:", *pizza)
}
