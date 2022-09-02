package src

import (
	"fmt"
)

type Loose struct {
	Gamer *Gamer
}

func (l *Loose) PrintResult() string {
	l.Gamer.Cash -= l.Gamer.Bet
	l.Gamer.Bet = l.Gamer.Cash / 10
	l.Gamer.SetState(l)
	//fmt.Printf("[ПРОИГРЫШ], на счету: %d, размер ставки: %d\n", l.Gamer.Cash, l.Gamer.Bet)
	return fmt.Sprintf("[ПРОИГРЫШ], на счету: %d, размер ставки: %d\n", l.Gamer.Cash, l.Gamer.Bet)
}
