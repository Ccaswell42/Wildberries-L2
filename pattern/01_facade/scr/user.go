package scr

// Структура User имеет карту, город, и имя
type User struct {
	Card *Card
	City City
	Name string
}

// конструктор
func NewUser(card *Card, City City, Name string) User {
	return User{Card: card, City: City, Name: Name}
}
