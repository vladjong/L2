package main

import "fmt"

/*
	Builder - интерфейс, который реализует поэтапное конструирование объекта
*/
type Builder interface {
	MakeName(val string)
	MakeColor(val string)
	MakeSize(val int)
}

/*
	Director - структура, которая управляет билдером
*/
type Director struct {
	builder Builder
}

/*
	ConstructGT - метод, который формирует структуру GT bike
*/
func (d *Director) ConstructGT() {
	d.builder.MakeName("GT")
	d.builder.MakeColor("Black")
	d.builder.MakeSize(17)
}

/*
	ConcreteBuilder - структура, которая реализует интерфейс Builder
*/
type ConcreteBuilder struct {
	product *Bike
}

func (b *ConcreteBuilder) MakeName(val string) {
	b.product.Name = val
}

func (b *ConcreteBuilder) MakeColor(val string) {
	b.product.Color = val
}

func (b *ConcreteBuilder) MakeSize(val int) {
	b.product.Size = val
}

/*
	Bike - абстрактный объект
*/
type Bike struct {
	Name  string
	Color string
	Size  int
}

func main() {
	obj := new(Bike)
	builder := ConcreteBuilder{
		product: obj,
	}
	director := Director{
		builder: &builder,
	}
	director.ConstructGT()
	fmt.Println(obj)
}
