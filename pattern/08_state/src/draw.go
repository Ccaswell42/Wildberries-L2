package src

import "fmt"

type Draw struct {
	Gamer *Gamer
}

// Roll устанавливает состояние объекта Gamer как "ничья"
func (d *Draw) Roll() {
	d.Gamer.SetState(d)
}

// PrintResult выводит информацию об объекте в состоянии "ничья"
func (d *Draw) PrintResult() {
	d.Gamer.SetState(d)
	fmt.Printf("[НИЧЬЯ], на счету: %d, размер ставки: %d\n", d.Gamer.Cash, d.Gamer.Bet)

}
