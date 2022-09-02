package src

import "fmt"

type Start struct {
	Gamer *Gamer
}

// Roll устанавливает состояние объекта Gamer как "старт игры"
func (s *Start) Roll() {
	s.Gamer.SetState(s)
}

// PrintResult выводит информацию об объекте в состоянии "старт игры"
func (s *Start) PrintResult() {
	fmt.Printf("[СТАРТ ИГРЫ], на счету: %d, размер ставки: %d\n", s.Gamer.Cash, s.Gamer.Bet)
}
