package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/
func main() {
	nums := []int{2, 4, 6, 8, 10}
	var sum int

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for _, value := range nums {
		wg.Add(1)
		go func(v int) {
			mutex.Lock()
			sum += v * v
			mutex.Unlock()
			wg.Done()
		}(value)
	}

	wg.Wait()
	fmt.Println(sum)
}
