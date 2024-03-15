package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// 2-ой способ. С помощью сигнального канала
func main() {
	ch := make(chan int, 1)    // буф. канал для передачи данных
	stop := make(chan bool, 1) // буф. канал для сигнала об остановке
	// запуск горутины
	go func(c chan int) {
		for {
			select { // если поступит сигнал в канал stop, то горутина завершит свою работу
			case <-stop:
				fmt.Println("goroutine stopped")
				return
			default: // дефолтная работа. вывод данных в консоль
				fmt.Println("val:", <-c)
			}
		}
	}(ch)

	// цикл с передачей данных в канал
	for i := 0; i < 5; i++ {
		ch <- i
	}
	// передаем в канал сигнал
	stop <- true
	// дожидаемся завершения горутины
	time.Sleep(time.Second)
}
