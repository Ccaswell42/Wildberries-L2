package src

type State interface {
	PrintResult() string
}

type Gamer struct {
	Bet   int
	Cash  int
	State State
}

func NewGamer(cash int) *Gamer {
	g := &Gamer{
		Bet:  cash / 10,
		Cash: cash,
	}
	st := &Start{g}
	return &Gamer{cash / 10, cash, st}
}

func (g *Gamer) SetState(st State) {
	g.State = st
}

func (g *Gamer) PrintResult() string {
	return g.State.PrintResult()

}
