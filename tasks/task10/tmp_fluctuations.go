package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc
*/
func main() {
	// объявляем последовательность чисел
	tmpFlict := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// объявляем мапу, где ключ - шаг в 10 градусов, а значение - слайс []float32
	seqMap := make(map[int][]float32)
	// проходим по слайсу со значениями
	for _, e := range tmpFlict {
		k := int(e/10) * 10              // получаем ключ - текущий диапазон, путем арифм. операций
		seqMap[k] = append(seqMap[k], e) // добавляем текущее значение в слайс по ключу
	}
	fmt.Println(seqMap) // выводим мапу
}
