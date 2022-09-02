package src

import "fmt"

type Loose struct {
	Gamer *Gamer
}

func (l *Loose) Roll() {
	l.Gamer.Cash -= l.Gamer.Bet
	l.Gamer.Bet = l.Gamer.Cash / 10
	l.Gamer.SetState(l)
}

func (l *Loose) PrintResult() {
	fmt.Printf("[ПРОИГРЫШ], на счету: %d, размер ставки: %d\n", l.Gamer.Cash, l.Gamer.Bet)

}
