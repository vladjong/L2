package main

import "fmt"

/*
	Шаблон проектирования State - основан на конечном автомате.
	Использовать нужно, когда:
	- объект может находиться в нескольких состояний. В зависимости от запроса он должен изменить свое состояние.
	- объект когда объект имеет разные ответы на один и тот же запрос в зависимости от состояния
	Предотвратит множественное if else

*/

type StateI interface {
	LightIsRed()
	LightIsYellow()
	LightIsGreen()
}

type TrafficLights struct {
	red          StateI
	green        StateI
	yellow       StateI
	currentState StateI
}

func (t *TrafficLights) SetState(state StateI) {
	t.currentState = state
}

func (t *TrafficLights) LightIsRed() {
	t.currentState.LightIsRed()
}

func (t *TrafficLights) LightIsYellow() {
	t.currentState.LightIsYellow()
}

func (t *TrafficLights) LightIsGreen() {
	t.currentState.LightIsGreen()
}

type Red struct {
	trafficLights *TrafficLights
}

func (r *Red) LightIsRed() {
	fmt.Println("The traffic light is red")
}

func (r *Red) LightIsYellow() {
	fmt.Println("Now traffic light is red")
	fmt.Println("Changing the color to yellow")
	r.trafficLights.SetState(r.trafficLights.yellow)
}

func (r *Red) LightIsGreen() {
	fmt.Println("The yellow color hasn't lit up yet")
}

type Yellow struct {
	trafficLights *TrafficLights
}

func (y *Yellow) LightIsRed() {
	fmt.Println("The green color hasn't lit up yet")
}

func (y *Yellow) LightIsYellow() {
	fmt.Println("The traffic light is yellow")
}

func (y *Yellow) LightIsGreen() {
	fmt.Println("Now traffic light is yellow")
	fmt.Println("Changing the color to green")
	y.trafficLights.SetState(y.trafficLights.green)
}

type Green struct {
	trafficLights *TrafficLights
}

func (g *Green) LightIsRed() {
	fmt.Println("Now traffic light is yellow")
	fmt.Println("Changing the color to green")
	g.trafficLights.SetState(g.trafficLights.green)
}

func (g *Green) LightIsYellow() {
	fmt.Println("The yellow color hasn't lit up yet")
}

func (g *Green) LightIsGreen() {
	fmt.Println("The traffic light is green")
}

func newTrafficLights() *TrafficLights {
	t := &TrafficLights{}
	red := &Red{
		trafficLights: t,
	}
	yellow := &Yellow{
		trafficLights: t,
	}
	green := &Green{
		trafficLights: t,
	}
	t.SetState(red)
	t.red = red
	t.yellow = yellow
	t.green = green
	return t
}

func main() {
	trafficLights := newTrafficLights()
	trafficLights.LightIsRed()
	trafficLights.LightIsYellow()
	trafficLights.LightIsGreen()
	trafficLights.LightIsYellow()
}
