package main

import (
	"state/src"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

func main() {
	bob := src.NewGamer(10000)
	bob.PrintResult()
	win := src.Win{Gamer: bob}
	lose := src.Loose{Gamer: bob}
	draw := src.Draw{Gamer: bob}

	bob.SetState(&win)
	bob.Roll()
	bob.PrintResult()
	bob.SetState(&win)
	bob.Roll()
	bob.PrintResult()
	bob.SetState(&draw)
	bob.Roll()
	bob.PrintResult()
	bob.SetState(&lose)
	bob.Roll()
	bob.PrintResult()
	bob.SetState(&lose)
	bob.Roll()
	bob.PrintResult()

}
