package main

import (
	"fmt"
)

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры
	Human (аналог наследования).
*/

/* структуру human*/
type human struct {
	Name    string
	Surname string
	age     int
}

// конструктор структуры human
func newHuman(name string, surname string, age int) *human {
	return &human{Name: name, Surname: surname, age: age}
}

// метод структуры human для проверки возраста
func (h human) isAdult() bool {
	if h.age < 18 {
		return false
	} else {
		return true
	}
}

// метод структуры human
func (h human) hello() {
	fmt.Printf("Hello, my name is %s\n", h.Name)
}

// структура action
type action struct {
	*human // встраиваем human
}

func main() {
	h := newHuman("Steve", "Stevenson", 21) // объявляем объект структуры human

	act := action{h} // объявляем объект структуры action

	// вызываем методы human у объекта структуры action
	fmt.Println(act.isAdult())
	act.hello()
}
