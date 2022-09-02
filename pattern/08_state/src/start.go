package src

import "fmt"

type Start struct {
	Gamer *Gamer
}

func (s *Start) PrintResult() string {
	s.Gamer.SetState(s)

	return fmt.Sprintf("[СТАРТ ИГРЫ], на счету: %d, размер ставки: %d\n", s.Gamer.Cash, s.Gamer.Bet)
}
