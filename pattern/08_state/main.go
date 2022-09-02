package main

import (
	"fmt"
	"state/src"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

func main() {
	bob := src.NewGamer(5000)
	fmt.Println(bob.PrintResult())
	//win := src.Win{Gamer: bob}
	lose := src.Loose{Gamer: bob}
	//draw := src.Draw{Gamer: bob}
	c := lose.PrintResult()
	fmt.Println(c)
	c = lose.PrintResult()
	fmt.Println(c)

	fmt.Println(bob.PrintResult())
}
