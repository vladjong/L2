package main

import "fmt"

type PizzaI interface {
	getPrice() int
}

type PeperonyPizza struct{}

func (p *PeperonyPizza) getPrice() int {
	return 10
}

type CheezePizza struct{}

func (c *CheezePizza) getPrice() int {
	return 15
}

type CheezeToping struct {
	pizza PizzaI
}

func (c *CheezeToping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 5
}

func main() {
	pepperonyPizze := &PeperonyPizza{}
	fmt.Println("Price peperony:", pepperonyPizze.getPrice())
	pepperonyPizzeWithCheese := &CheezeToping{
		pizza: pepperonyPizze,
	}
	fmt.Println("Price peperony with cheese:", pepperonyPizzeWithCheese.getPrice())
}
