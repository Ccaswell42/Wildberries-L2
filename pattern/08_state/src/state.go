package src

// State предоставляет интерфейс для методов состояний.
type State interface {
	PrintResult()
	Roll()
}

// Gamer - структура игрока
type Gamer struct {
	Bet   int
	Cash  int
	State State
}

// NewGamer - конструктор
func NewGamer(cash int) *Gamer {
	g := &Gamer{
		Bet:  cash / 10,
		Cash: cash,
	}
	st := &Start{g}
	return &Gamer{cash / 10, cash, st}
}

// SetState - меняет состояние объекта
func (g *Gamer) SetState(st State) {
	g.State = st
}

// Roll - реализуетп поведение текущего состояния
func (g *Gamer) Roll() {
	g.State.Roll()
}

// PrintResult выводит информацию об объекте в текущем состоянии
func (g *Gamer) PrintResult() {
	g.State.PrintResult()

}
