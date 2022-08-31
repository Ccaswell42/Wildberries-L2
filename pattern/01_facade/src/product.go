package src

// переменные для примера работы магазина
var (
	HairDryer    = Product{"Dyson", 35000}
	TV           = Product{"Samsung", 50000}
	Laptop       = Product{"Apple", 90000}
	Sunglasses   = Product{"Ray Ban", 15000}
	Moscow       = City{"Moscow"}
	StPetersburg = City{"StPetersburg"}
	Kazan        = City{"Kazan"}
	Samara       = City{"Samara"}
)

// структура Product, имитирующая товар, имеет название товара и его цену.
type Product struct {
	Name  string
	Price int
}

// структура City хранит название города
type City struct {
	Name string
}

// Конструктор
func NewCity(name string) City {
	return City{name}
}
