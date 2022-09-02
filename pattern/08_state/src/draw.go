package src

import "fmt"

type Draw struct {
	Gamer *Gamer
}

func (d *Draw) PrintResult() string {
	d.Gamer.SetState(d)
	//fmt.Printf("[НИЧЬЯ], на счету: %d, размер ставки: %d\n", d.Gamer.Cash, d.Gamer.Bet)
	return fmt.Sprintf("[НИЧЬЯ], на счету: %d, размер ставки: %d\n", d.Gamer.Cash, d.Gamer.Bet)
}
