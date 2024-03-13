package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

// N - секунд после чего программа завершиться
const timeToWork = 3

func main() {
	stop := make(chan bool, 1) // буф. канал для сообщения горутине об остановке

	dataChan := make(chan int, 1) // основной буф. канал для передачи данных
	defer close(dataChan)         // заранее объявляем о закрытии канала

	timer := time.After(timeToWork * time.Second) // таймер после которого программа должна завершиться

	// Вызов горутины
	go func(dc chan int) {
		for {
			select {
			// завершение горутины при получении значения в канал stop
			case <-stop:
				fmt.Printf("Stop reader goroutine")
				return
			case val := <-dc: // считывания значения из канала и вывод его в консоль
				fmt.Printf("value=%d\n", val)
			}
		}
	}(dataChan)

	// основной цикл main
	for {
		select {
		case <-timer: // по истечении таймера, в канал поступит сигнал и программа завершиться
			fmt.Println("Stop writing")
			stop <- true            // подаем сигнал в канал остановки горутины
			time.Sleep(time.Second) // ожидаем завершения читающей горутины
			return
		default: // дефолтная работа. Запись рандомных значений в канал
			dataChan <- rand.Int()
			time.Sleep(time.Second / 2)
		}
	}
}
