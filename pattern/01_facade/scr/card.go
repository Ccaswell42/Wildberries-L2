package scr

import "errors"

// Структура Card хранит баланс карты и название банка, выпустившего карту.
type Card struct {
	BankName string
	Balance  int
}

// конструктор
func NewCard(balance int) *Card {
	return &Card{BankName: "Sber", Balance: balance}
}

// функция списания денег с баланса карты
func (c *Card) WithdrawRequest(sum int) error {
	if c.Balance < sum {
		return errors.New("недостаточно средств на карте")
	}
	c.Balance -= sum
	return nil
}
