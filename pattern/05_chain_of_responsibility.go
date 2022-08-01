package main

/*
	Паттерн Chain of responsible - это поведенческий шаблон проектирования, позволяет создавать цепочку обработчиков запросов.
	Для каждого входящего запроса он передается по цепочке и каждому из обработчиков.
	Когда использовать:
	- шаблон применим, когда есть несколько кандидатов для обработки одного и того же запроса.
	- когда клиент не знает получателя, ему не надо выбирать получателя
*/

import "fmt"

/*
	Структура клиента
*/
type Intern struct {
	name                string
	jobPosition         string
	registrationDone    bool
	interviewWithHRDone bool
	finalInterview      bool
}

/*
	Интерфейс, который будут имплементировать методы цепочки
*/
type Intership interface {
	Execute(*Intern)
	SetNext(Intership)
}

type Application struct {
	next Intership
}

func (a *Application) Execute(intern *Intern) {
	if intern.registrationDone {
		fmt.Println("Intern has already filled out the questionnaire!")
		a.next.Execute(intern)
		return
	}
	fmt.Println("Intern filled out the questionnaire!")
	intern.registrationDone = true
	a.next.Execute(intern)
}

func (a *Application) SetNext(next Intership) {
	a.next = next
}

type HR struct {
	next Intership
}

func (h *HR) Execute(intern *Intern) {
	if intern.interviewWithHRDone {
		fmt.Println("Intern has already been interviewed by HR!")
		h.next.Execute(intern)
		return
	}
	fmt.Println("Intern was interviewed by HR!")
	intern.interviewWithHRDone = true
	h.next.Execute(intern)
}

func (h *HR) SetNext(next Intership) {
	h.next = next
}

type Teamlead struct {
	next Intership
}

func (t *Teamlead) Execute(intern *Intern) {
	if intern.finalInterview {
		fmt.Println("Intern has already been interviewed by Teamlead!")
	}
	fmt.Println("Intern was interviewed by Teamlead!")
}

func (t *Teamlead) SetNext(next Intership) {
	t.next = next
}

func main() {
	teamlead := &Teamlead{}
	hr := &HR{}
	hr.SetNext(teamlead)
	application := &Application{}
	application.SetNext(hr)
	intern := &Intern{
		name:        "Vlad",
		jobPosition: "Backend develop",
	}
	fmt.Println("The process of employment of the intern:", intern.name)
	fmt.Println("To the position:", intern.jobPosition)
	application.Execute(intern)
}
