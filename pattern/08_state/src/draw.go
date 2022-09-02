package src

import "fmt"

type Draw struct {
	Gamer *Gamer
}

func (d *Draw) Roll() {
	d.Gamer.SetState(d)
}

func (d *Draw) PrintResult() {
	d.Gamer.SetState(d)
	fmt.Printf("[НИЧЬЯ], на счету: %d, размер ставки: %d\n", d.Gamer.Cash, d.Gamer.Bet)

}
