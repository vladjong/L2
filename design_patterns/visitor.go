package main

/*
	Паттерн Visitor - расширяет функциональность класса, не изменяя его первоначальную реализацию
*/

import "fmt"

/*
	Интерфейс Visitor, в котором содержится метод определения года выпуска модели
*/
type Visitor interface {
	VisitCarYear(*Car)
}

/*
	Структура Car, который содержит поле name
*/
type Car struct {
	name string
}

/*
	Метод Accept, который добавляет функционал посетителя в наш класс
*/
func (car *Car) Accept(visitor Visitor) {
	visitor.VisitCarYear(car)
}

/*
	Структура RealeseCalendar, которая содержит поле year
*/
type ReleaseCalendar struct {
	year int
}

/*
	Имплементация метода интерфейса Visitor структурой RealeseCalendar
*/
func (release *ReleaseCalendar) VisitCarYear(car *Car) {
	if car.name == "BMW X6" {
		fmt.Println("You car BMW X6")
		release.year = 2010
		return
	}
	release.year = 2020
}

func main() {
	car := new(Car)
	car.name = "BMW X6"
	release := new(ReleaseCalendar)
	car.Accept(release)
	fmt.Println(release.year)
}
