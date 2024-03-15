package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map
*/

// структура мапы с мьютексом
type MuMap struct {
	values map[int]int
	sync.Mutex
}

// метод MuMap для конкурентной записи
func (m *MuMap) add(k, v int) {
	m.Lock()
	m.values[k] = v
	m.Unlock()
}

func main() {
	nums := MuMap{values: make(map[int]int)} // объявляем объект структуры MuMap
	var wg sync.WaitGroup                    // объявляем переменную WaitGroup для синхронизации горутин

	// цикл для записи в мапу
	for i := 0; i < 10; i++ {
		wg.Add(1) // увеличиваем счетчик на 1
		// запуск горутины
		go func(n *MuMap, val int) {
			n.add(val, val)
			wg.Done() // уменьшаем счетчик на 1
		}(&nums, i)
	}

	// Дожидаемся завершения всех горутин
	wg.Wait()
	// выводим значения
	fmt.Println(nums.values)
}
