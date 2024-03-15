package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// структуры которые адаптируем
type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("bark, bark!!!")
}

type Worker struct {
}

func (w Worker) Work() {
	fmt.Println("Im do some work")
}

// target - Интерфейс адаптера
type ActionAdapter interface {
	Action()
}

// Адаптер собаки
type DogAdapter struct {
	*Dog
}

func NewDogAdapter() *DogAdapter {
	return &DogAdapter{Dog: &Dog{}}
}

func (d DogAdapter) Action() {
	d.Bark()
}

// адаптер рабочего
type WorkerAdapter struct {
	*Worker
}

func NewWorkerAdapter() *WorkerAdapter {
	return &WorkerAdapter{Worker: &Worker{}}
}

func (w WorkerAdapter) Action() {
	w.Work()
}

// У студента уже реализован метод, поэтому его адаптировать не нужно
type Student struct {
}

func (s Student) Action() {
	fmt.Println("Im learning")
}

// пример использования
func main() {
	someActioners := []ActionAdapter{NewDogAdapter(), NewWorkerAdapter(), Student{}}
	for _, e := range someActioners {
		doAction(e)
	}
}

// функция принимающая адаптер
func doAction(actioner ActionAdapter) {
	actioner.Action()
}
