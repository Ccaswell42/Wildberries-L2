package main

import "fmt"

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Visitor_pattern

*/

/*

 */
/*

 */
/*

 */

type Visitor interface {
	visitStore(place Place) string
	visitPharmacy(place Place) string
	visitBarbershop(place Place) string
}
type Citizen struct{}

func (c *Citizen) visitStore(place Place) string {
	fmt.Printf("Pay for furniture in %T, %s\n", place, place.Print())
	return fmt.Sprintf("Pay for furniture in %T\n", place)
}
func (c *Citizen) visitPharmacy(place Place) string {
	fmt.Printf("Pay for drugs in %T, %s\n", place, place.Print())
	return fmt.Sprintf("Pay for drugs in %T\n", place)
}
func (c *Citizen) visitBarbershop(place Place) string {
	fmt.Printf("Pay for haircut in %T, %s\n", place, place.Print())
	return fmt.Sprintf("Pay for haircut in %T\n", place)
}

type Place interface {
	Accept(v Visitor) string
	Print() string
}

type Store struct {
	Name string
}

func (s *Store) Print() string {
	return s.Name
}

func (s *Store) Accept(v Visitor) string {
	return v.visitStore(s)
}

type Pharmacy struct {
	Name string
}

func (p *Pharmacy) Print() string {
	return p.Name
}

func (p *Pharmacy) Accept(v Visitor) string {
	return v.visitPharmacy(p)
}

type Barbershop struct {
	Name string
}

func (b *Barbershop) Print() string {
	return b.Name
}

func (b *Barbershop) Accept(v Visitor) string {
	return v.visitBarbershop(b)
}

type AllMarkets struct {
	b *Barbershop
	s *Store
	p *Pharmacy
}

func NewAllMarkets() *AllMarkets {
	return &AllMarkets{
		&Barbershop{Name: "Chop-chop"},
		&Store{Name: "Hoff"},
		&Pharmacy{Name: "Dr. Aibolit"},
	}
}
func (all *AllMarkets) Visit(v Visitor) {
	v.visitStore(all.s)
	v.visitPharmacy(all.p)
	v.visitBarbershop(all.b)
}

func main() {
	all := NewAllMarkets()
	user := Citizen{}
	all.Visit(&user)
}
