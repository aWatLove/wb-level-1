package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

// 1-ый способ. С использованием мьютекса

type Count struct {
	count int
	sync.Mutex
}

func (c *Count) increment() {
	c.Lock()
	c.count++
	c.Unlock()
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
