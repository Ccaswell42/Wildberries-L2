package src

import "fmt"

type Loose struct {
	Gamer *Gamer
}

// Roll устанавливает состояние объекта Gamer как "проигрыш" и реализует поведение проигрыша ставки
func (l *Loose) Roll() {
	l.Gamer.Cash -= l.Gamer.Bet
	l.Gamer.Bet = l.Gamer.Cash / 10
	l.Gamer.SetState(l)
}

// PrintResult выводит информацию об объекте в состоянии "проигрыш"
func (l *Loose) PrintResult() {
	fmt.Printf("[ПРОИГРЫШ], на счету: %d, размер ставки: %d\n", l.Gamer.Cash, l.Gamer.Bet)

}
