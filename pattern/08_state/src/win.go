package src

import "fmt"

type Win struct {
	Gamer *Gamer
}

func (w *Win) PrintResult() string {
	fmt.Println(w.Gamer.Bet)
	w.Gamer.Cash += w.Gamer.Bet
	w.Gamer.Bet = w.Gamer.Cash / 10
	w.Gamer.SetState(w)
	//fmt.Printf("[ПОБЕДА], на счету: %d, размер ставки: %d\n", w.Gamer.Cash, w.Gamer.Bet)
	return fmt.Sprintf("[ПОБЕДА], на счету: %d, размер ставки: %d\n", w.Gamer.Cash, w.Gamer.Bet)
}
