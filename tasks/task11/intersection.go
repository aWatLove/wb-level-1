package main

import "fmt"

func main() {
	// объявляем два слайса
	list1 := []int{1, 2, 7, 4, 9}
	list2 := []int{7, 5, 2, 11, 10, 3, 5, 4}

	// мапа, где ключ - значение из наименьшего слайса, а значение - пустая структура
	inter := make(map[int]struct{})

	// сравниваем длину и при надобности меняем ссылки
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}

	// записываем в мапу значения из наименьшего слайса
	for _, e := range list1 {
		inter[e] = struct{}{}
	}

	// результрующий слайс
	res := make([]int, 0)

	// цикл по наибольшему слайсу
	for _, e := range list2 {
		_, ok := inter[e] // если ok == true, то записываем в результрующий слайс
		if ok {
			res = append(res, e)
		}
	}
	// выводим результат
	fmt.Println(res)
}
