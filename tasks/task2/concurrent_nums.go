package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
func main() {
	nums := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	for _, value := range nums {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(v * v)
		}(value)
	}

	wg.Wait()
}
