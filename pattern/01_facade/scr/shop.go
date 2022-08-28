package scr

import (
	"errors"
	"fmt"
	"time"
)

// Наш магащин имеет какие-то товары дла продажи, и базу данных городов, в которые может доставлять товары
type Shop struct {
	Products     []Product
	CityDelivery []City
}

// для примера заполним базу данных городов нашего магазина
func (shop *Shop) SetCityDeliverDB() {
	shop.CityDelivery = append(shop.CityDelivery, Moscow, Kazan, Samara, StPetersburg)
}

// для примера заполним наш магазин некоторыми товарами
func (shop *Shop) SetProductsDB() {
	shop.Products = append(shop.Products, HairDryer, Laptop, TV, Sunglasses)
}

// Конструируем наш магазин сразу с товарами и базой городов
func NewShop() *Shop {
	var shop Shop
	shop.SetProductsDB()
	shop.SetCityDeliverDB()
	return &shop
}

// Основная функция магазина - продажа товара.
//Функция Sell принимает пользователя и товар который он хочет купить.

func (shop Shop) Sell(user User, product string) error {
	var ourProduct Product
	fmt.Println("Проверка наличия товара на складе")
	time.Sleep(500 * time.Millisecond)
	for i, val := range shop.Products {
		if product == val.Name {
			ourProduct = val
			break
		}
		if i == len(shop.Products)-1 {

			return errors.New("такого товара нет на складе")
		}
	}

	fmt.Println("Проверка баланса для покупки товара")
	time.Sleep(500 * time.Millisecond)
	if user.Card.Balance < ourProduct.Price {
		return errors.New("недостаточно средств на карте")
	}
	fmt.Println("Проверка возможности доставки в данный город")
	time.Sleep(500 * time.Millisecond)
	for i, val := range shop.CityDelivery {
		if user.City.Name == val.Name {
			break
		}
		if i == len(shop.Products)-1 {
			return errors.New("доставка в данный город недоступна")
		}
	}
	err := user.Card.WithdrawRequest(ourProduct.Price) // <- списываем деньги со счета клиента
	if err != nil {
		return err
	}
	fmt.Println("Покупка совершена успешно")
	fmt.Printf("Товар %s скоро будет доставлен в %s, остаток по карте: %d\n", ourProduct.Name, user.City.Name, user.Card.Balance)
	return nil
}

// наш фасад, принимает город доставки, необходимы товар, имя заказчика и количество денег на счету заказчика
// из полученных данных конструируем нужные объекты и реализуем работу наешго магазина.

func Facade(city, product, name string, money int) error {
	card := NewCard(money)
	newCity := NewCity(city)
	user := NewUser(card, newCity, name)
	shop := NewShop()
	err := shop.Sell(user, product)
	if err != nil {
		return err
	}
	return nil
}
