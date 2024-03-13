package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// 3-ий способ. С помощью контекста
func main() {
	ch := make(chan int, 1) // буф. канал для передачи данных

	ctx, cancel := context.WithCancel(context.Background()) // объявляем контекст с функцией отмены
	// запуск горутины
	go func(c chan int) {
		for {
			select { // если контекст отменен, выбирается этот случай
			case <-ctx.Done():
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
	// вызываем функцию cancel, чтобы отменить контекст
	cancel()
	// дожидаемся завершения горутины
	time.Sleep(time.Second)
}
