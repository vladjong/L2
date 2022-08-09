package main

import "fmt"

/*
	Factory design паттерн - это шаблон позволяет скрыть логику создания экземпляров.
	Клиент взаимодействует только с factory struct и сообщает тип экземпляра, которые необходимо создать.
*/

type AppleI interface {
	SetName(name string)
	SetCategory(category string)
	GetName() string
	GetCategory() string
}

type device struct {
	name     string
	category string
}

func (d *device) SetName(name string) {
	d.name = name
}

func (d *device) SetCategory(category string) {
	d.category = category
}

func (d *device) GetName() string {
	return d.name
}

func (d *device) GetCategory() string {
	return d.category
}

type Iphone struct {
	device
}

func newIphone() AppleI {
	return &Iphone{
		device: device{
			name:     "Iphone 13",
			category: "Phone",
		},
	}
}

type MacBook struct {
	device
}

func newMacBook() AppleI {
	return &MacBook{
		device: device{
			name:     "MacBook pro 14",
			category: "Notebook",
		},
	}
}

func getDevice(device string) (AppleI, error) {
	if device == "iphone" {
		return newIphone(), nil
	} else if device == "macbook" {
		return newMacBook(), nil
	}
	return nil, fmt.Errorf("Wrong device type")
}

func printDetails(d AppleI) {
	fmt.Println("Device:", d.GetName())
	fmt.Println("Category:", d.GetCategory())
}

func main() {
	iphone, _ := getDevice("iphone")
	macbook, _ := getDevice("macbook")
	printDetails(iphone)
	printDetails(macbook)
}
