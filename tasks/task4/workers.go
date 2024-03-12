package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

// функция, которая будет запускаться в качестве горутины
func worker(id int, ch <-chan int, stop <-chan os.Signal) {
	for {
		select {
		case <-stop:
			fmt.Printf("stop working goroutine with id=%d\n", id)
			return
		case val := <-ch:
			fmt.Printf("goroutine id=%d value=%d\n", id, val)
		}
	}
}

// Кол-во горутин
const workerAmount = 5

func main() {
	// Объявление буферизированного канала для данных
	// размер буфера равен кол-ву горутин
	ch := make(chan int, workerAmount)
	defer close(ch) // объявляем об отложенном закрытии канала

	mainSignCh := make(chan os.Signal, 1) // буферизированный канал для сигнала

	// Объявляем какие сигналы "слушать" методом Notify, чтобы передать об этом в наш канал mainSignCh
	signal.Notify(mainSignCh, syscall.SIGTERM, syscall.SIGINT)

	// Запуск горутин в цикле,
	// а также создание сигнального канала для каждой горутины
	for i := 0; i < workerAmount; i++ {
		signCh := make(chan os.Signal, 1)
		signal.Notify(signCh, syscall.SIGTERM, syscall.SIGINT)

		go worker(i, ch, signCh)
	}

	// основной цикл main
	for {
		// по-умолчанию записываются новые данные в канал, либо при получении сигнала, поток останавливается
		select {
		case <-mainSignCh:
			fmt.Println("main is stopped")
			return
		default:
			ch <- rand.Int()
			time.Sleep(time.Millisecond * 100)
		}

	}
}
