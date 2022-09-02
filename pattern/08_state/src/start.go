package src

import "fmt"

type Start struct {
	Gamer *Gamer
}

func (s *Start) Roll() {
	s.Gamer.SetState(s)
}
func (s *Start) PrintResult() {
	fmt.Printf("[СТАРТ ИГРЫ], на счету: %d, размер ставки: %d\n", s.Gamer.Cash, s.Gamer.Bet)
}
