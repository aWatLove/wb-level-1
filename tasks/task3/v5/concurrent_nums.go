package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/
func main() {
	nums := []int32{2, 4, 6, 8, 10}
	var sum int32

	wg := sync.WaitGroup{}
	for _, value := range nums {
		wg.Add(1)
		go func(v int32) {
			atomic.AddInt32(&sum, v*v)
			wg.Done()
		}(value)
	}

	wg.Wait()

	fmt.Println(sum)
}
