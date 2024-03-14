package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

// 1-ый способ. С использованием атомиков

type Count struct {
	count int64
}

func (c *Count) increment() {
	atomic.AddInt64(&c.count, 1)
}

func main() {
	c := Count{}
	var wg sync.WaitGroup

	// в цикле запускаем 10 горутин, где каждая горутина инкрементирует счетчик 10 раз
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				c.increment()
			}
			wg.Done()
		}()
	}
	// дожидаемся завершения горутин
	wg.Wait()
	fmt.Println(c.count)
}
