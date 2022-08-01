package main

import "fmt"

/*
	Strategy паттерн - это шаблон позволяет изменять поведение объекта во время выполнения программы.
*/

type StrategyProjection interface {
	projection()
}

type CentralProjection struct{}

func (c *CentralProjection) projection() {
	fmt.Println("The construction is performed in the central projection")
}

type ParallelProjection struct{}

func (p *ParallelProjection) projection() {
	fmt.Println("The construction is performed in a parallel projection")
}

type Model struct {
	name     string
	strategy StrategyProjection
}

func NewModel(name string, strategy StrategyProjection) *Model {
	return &Model{
		name:     name,
		strategy: strategy,
	}
}

func (m *Model) SetStrategy(strategy StrategyProjection) {
	m.strategy = strategy
}

func (m *Model) Draw() {
	m.strategy.projection()
}

func main() {
	central := &CentralProjection{}
	parallel := &ParallelProjection{}
	model := NewModel("triangle", central)
	model.Draw()
	model.SetStrategy(parallel)
	model.Draw()
}
