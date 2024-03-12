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

	for _, value := range nums {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			sum += v * v
		}(value)
	}

	wg.Wait()
	fmt.Println(sum)
}
