package main

import "fmt"

/*
	Паттерн Command позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект. Такие запросы, например, можно ставить в очередь, отменять или возобновлять.
*/

/*
	Command предоставляет командный интерфейс
*/
type Command interface {
	Execute() string
}

/*
	ToggleXCommand реализует Comand
*/
type ToggleXCommand struct {
	receiver *Receiver
}

func (c *ToggleXCommand) Execute() string {
	return c.receiver.ToggleX()
}

/*
	ToggleYCommand реализует Comand
*/
type ToggleYCommand struct {
	receiver *Receiver
}

func (c *ToggleYCommand) Execute() string {
	return c.receiver.ToggleY()
}

/*
	ToggleZCommand реализует Comand
*/
type ToggleZCommand struct {
	receiver *Receiver
}

func (c *ToggleZCommand) Execute() string {
	return c.receiver.ToggleZ()
}

/*
	ToggleZCommand реализует Comand
*/
type ToggleOCommand struct {
	receiver *Receiver
}

func (c *ToggleOCommand) Execute() string {
	return c.receiver.ToggleO()
}

/*
	Реализация Receiver
*/
type Receiver struct {
}

func (r *Receiver) ToggleX() string {
	return "Give a pass"
}

func (r *Receiver) ToggleY() string {
	return "Give a pass in the cut"
}

func (r *Receiver) ToggleO() string {
	return "Hit the goal"
}

func (r *Receiver) ToggleZ() string {
	return "Take away the ball"
}

/*
	Реализация Invoker, который хранит в себе слайс Comand
*/
type Invoker struct {
	commands []Command
}

/*
	Добавление команд
*/
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

/*
	Удаление команды
*/
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

/*
	Метод, который выводит все команды доступные пользователю
*/
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}

func main() {
	receiver := new(Receiver)
	x := ToggleXCommand{receiver: receiver}
	y := ToggleYCommand{receiver: receiver}
	z := ToggleZCommand{receiver: receiver}
	o := ToggleOCommand{receiver: receiver}
	invoker := new(Invoker)
	invoker.StoreCommand(&x)
	invoker.StoreCommand(&y)
	invoker.StoreCommand(&z)
	invoker.StoreCommand(&o)
	invoker.StoreCommand(&o)
	invoker.UnStoreCommand()
	fmt.Println("The command menu for the Fifa game:")
	fmt.Print(invoker.Execute())
}
