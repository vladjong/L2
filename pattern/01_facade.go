package main

/*
	Шаблон проектирования Фасад - это структура, которая маскирует более сложную систему и служит интерфейсом для пользователя.
*/

import (
	"errors"
	"fmt"
	"time"
)

type Product struct {
	Name string
	Cost float64
}

/*
	Структура shop является фасадом
*/
type Shop struct {
	Name     string
	Products []Product
}

/*
	Метод Sell скрывает реализацию сложной системы [покупки товара]
*/
func (s *Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю, для получение остатка")
	time.Sleep(time.Second)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Println("[Магазин] Проверка - может ли пользователь", user.Name, "купить товар!")
	for _, prod := range s.Products {
		if prod.Name == product {
			if prod.Cost > user.GetBalance() {
				return errors.New("[Магазин] Не достаточно средств!")
			}
		}
	}
	fmt.Println("[Магазин] Пользователь купил товар", product)
	return nil
}

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (c *Card) CheckBalance() error {
	fmt.Println("[Карта] запрос в банк для проверки остатка")
	time.Sleep(time.Second)
	return c.Bank.CheckBalance(c.Name)
}

type Bank struct {
	Name  string
	Cards []Card
}

func (b *Bank) CheckBalance(cardNumber string) error {
	fmt.Println("[Банка] получение остатка по карте", cardNumber)
	time.Sleep(time.Second)
	for _, card := range b.Cards {
		if card.Name == cardNumber {
			if card.Balance < 0 {
				return errors.New("[Банк] Не достаточно средств!")
			}
		}
	}
	fmt.Println("[Банк] Остаток положительный!")
	return nil
}

type User struct {
	Name string
	Card *Card
}

func (u *User) GetBalance() float64 {
	return u.Card.Balance
}

var (
	bank = Bank{
		Name:  "Сбербанк",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "CRD-1",
		Balance: 150,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "CRD-2",
		Balance: 2,
		Bank:    &bank,
	}
	user1 = User{
		Name: "U-1",
		Card: &card1,
	}
	user2 = User{
		Name: "U-2",
		Card: &card2,
	}
	prod1 = Product{
		Name: "Water",
		Cost: 10,
	}
	prod2 = Product{
		Name: "Ring",
		Cost: 1000,
	}
	shop = Shop{
		Name:     "Lenta",
		Products: []Product{prod1, prod2},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск карт!")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Println(user1.Name)
	err := shop.Sell(user1, prod1.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user2.Name)
	err = shop.Sell(user2, prod1.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user1.Name)
	err = shop.Sell(user1, prod2.Name)
	if err != nil {
		fmt.Println(err)
	}
}
