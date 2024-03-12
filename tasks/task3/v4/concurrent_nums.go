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
	ch := make(chan int, 5)

	wg := sync.WaitGroup{}
	go func() {
		defer close(ch)
		for _, value := range nums {
			wg.Add(1)
			go func(v int) {
				ch <- v * v
				wg.Done()
			}(value)
		}
		wg.Wait()
	}()

	for val := range ch {
		sum += val
	}
	fmt.Println(sum)
}
