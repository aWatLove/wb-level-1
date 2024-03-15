package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

// 1-ый способ. С помощью time.After
func sleep(d time.Duration) {
	if <-time.After(d); true {
		return
	}
}

// 2-ой способ. С помощью таймера
func sleep2(d time.Duration) {
	timer := time.NewTimer(d)
	if <-timer.C; true {
		return
	}
}

func main() {
	duration := time.Second * 5 // время ожидания
	fmt.Println(time.Now())
	sleep(duration) // 1-я функция
	fmt.Println(time.Now())
	sleep2(duration) // 2-я функция
	fmt.Println(time.Now())
}
