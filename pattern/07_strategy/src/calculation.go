package src

// Calculation обеспечивает контекст для выполнения стратегии.
type Calculation struct {
	Strategy
}

// SetCalculation заменяет стратегии.
func (c *Calculation) SetCalculation(str Strategy) {
	c.Strategy = str
}
