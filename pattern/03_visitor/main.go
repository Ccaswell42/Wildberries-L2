package main

import "fmt"

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Visitor_pattern

*/

/*
Паттерн посетитель - поведенческий шаблон проектирования, позволяющий добавить новые функции имеющимся объектам,
без вмешательства в код объектов

*/

/*
Плюсы:
	1. Упрощает добавление новых операций.
	2. Объединяет родственные операции в одном классе.
	3. Посетитель может накапливать состояние при обходе структуры элементов.
*/

/*
Недостатки:
Затрудняет расширение иерархии классов, поскольку новые классы обычно требуют добавления нового
visit метода для каждого посетителя.
*/

/*
Реализуем паттерн посетитель на прримере похода человека в магазин, аптеку и парикмахерскую.
*/

// интерфейс посетителя
type Visitor interface {
	visitStore(place *Store) string
	visitPharmacy(place *Pharmacy) string
	visitBarbershop(place *Barbershop) string
}

// структура посетителя, реализующая интерфейс Visitor
type Citizen struct{}

//Методы новой функциональности для наших объектов:

// посетитель посещает объект типа "магазин" и совершает покупку мебели
func (c *Citizen) visitStore(place *Store) string {
	fmt.Printf("Pay for furniture in %T, %s\n", place, place.Print())
	return fmt.Sprintf("Pay for furniture in %T\n", place)
}

// посетитель посещает бъект типа "аптека" и совершает покупку лекартсв
func (c *Citizen) visitPharmacy(place *Pharmacy) string {
	fmt.Printf("Pay for drugs in %T, %s\n", place, place.Print())
	return fmt.Sprintf("Pay for drugs in %T\n", place)
}

// посеитель посещает объект типа "барбершоп" и платит за услугу
func (c *Citizen) visitBarbershop(place *Barbershop) string {
	fmt.Printf("Pay for haircut in %T, %s\n", place, place.Print())
	return fmt.Sprintf("Pay for haircut in %T\n", place)
}

// наша структура с возможностью приема посетителей

type AllMarkets struct {
	b *Barbershop
	s *Store
	p *Pharmacy
}

// конструктор
func NewAllMarkets() *AllMarkets {
	return &AllMarkets{
		&Barbershop{Name: "Chop-chop"},
		&Store{Name: "Hoff"},
		&Pharmacy{Name: "Dr. Aibolit"},
	}
}

// структура магазина
type Store struct {
	Name string
}

func (s *Store) Print() string {
	return s.Name
}

// структура аптеки
type Pharmacy struct {
	Name string
}

func (p *Pharmacy) Print() string {
	return p.Name
}

// структура барбершопа
type Barbershop struct {
	Name string
}

func (b *Barbershop) Print() string {
	return b.Name
}

// Принимает посетителей и реализукт новую функциональность для наших объектов

func (all *AllMarkets) Visit(v Visitor) {
	v.visitStore(all.s)
	v.visitPharmacy(all.p)
	v.visitBarbershop(all.b)
}

func main() {
	all := NewAllMarkets() // инициализируем нашу структуру
	user := Citizen{}      // создаем посетителя
	all.Visit(&user)       // посетитель реализует новые функции
}
