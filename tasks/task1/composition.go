package main

import (
	"fmt"
)

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры
	Human (аналог наследования).
*/
/* Объявляем структуру human*/
type human struct {
	Name    string
	Surname string
	age     int
}

func newHuman(name string, surname string, age int) *human {
	return &human{Name: name, Surname: surname, age: age}
}

func (h human) isAdult() bool {
	if h.age < 18 {
		return false
	} else {
		return true
	}
}

func (h human) hello() {
	fmt.Printf("Hello, my name is %s\n", h.Name)
}

// объявляем структуру action
type action struct {
	*human // встраиваем human
}

func main() {
	h := newHuman("Steve", "Stevenson", 21)
	act := action{h}
	fmt.Println(act.isAdult())
	act.hello()
}
