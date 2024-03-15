package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// 1-ый способ. С помощью закрытия канала
func main() {
	ch := make(chan int, 1) // буф. канал для передачи данных

	// запуск горутины
	go func(c chan int) {
		for {
			val, ok := <-c
			// если канал закрыт, ok == false, то завершаем работу горутины
			if !ok {
				fmt.Println("Goroutine stopped")
				return
			}
			fmt.Println("value:", val)
		}
	}(ch)

	// цикл с передачей данных в канал
	for i := 0; i < 5; i++ {
		ch <- i
	}
	// закрываем канал
	close(ch)
	// дожидаемся завершения горутины
	time.Sleep(time.Second)
}
