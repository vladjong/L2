package main

import "fmt"

type ReceptI interface {
	addDrink()
}

type Starbucks struct {
	reseptI ReseptI
}

func (s *Starbucks) makeTemplateOrder() {
	s.boilWater()
	s.reseptI.addDrink()
	s.pointIntoGlass()
}

func (s *Starbucks) boilWater() {
	fmt.Println("Boil water")
}

func (s *Starbucks) pointIntoGlass() {
	fmt.Println("Pour the drink into a glass")
}

type Coffee struct {
	Starbucks
}

func (c *Coffee) addDrink() {
	fmt.Println("Add coffee to a drink")
}

type Tea struct {
	Starbucks
}

func (t *Tea) addDrink() {
	fmt.Println("Add tea to a drink")
}

func main() {
	coffee := &Coffee{}
	starbucks := Starbucks{
		reseptI: coffee,
	}
	starbucks.makeTemplateOrder()
}
