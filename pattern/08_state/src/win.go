package src

import "fmt"

type Win struct {
	Gamer *Gamer
}

func (w *Win) Roll() {
	w.Gamer.Cash += w.Gamer.Bet
	w.Gamer.Bet = w.Gamer.Cash / 10
	w.Gamer.SetState(w)
}

func (w *Win) PrintResult() {
	fmt.Printf("[ПОБЕДА], на счету: %d, размер ставки: %d\n", w.Gamer.Cash, w.Gamer.Bet)

}
