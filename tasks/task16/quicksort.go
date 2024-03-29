package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/
func quicksort(arr []int, begin, end int) {
	// если границы сомкнулись
	if end-begin <= 1 {
		return
	}

	i := begin
	j := end
	q := arr[begin+(end-begin)/2] // опорный элемент

	for i <= j {
		// икрементируем индекс левой границы пока текущий элемент меньше опорного
		for arr[i] < q {
			i++
		}
		// декрементируем индекс правой границы пока элемент больше опорного
		for arr[j] > q {
			j--
		}

		// если индекс левой границы меньше правой, то меняем элементы местами
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++ // увеличиваем индекс левой границы
			if j > 0 {
				j--
			}
		}
	}

	// если левая граница меньше опорного, то рекурсивно вызываем функцию сортировки
	// передавая указатель на слайс, текущую левую границу и новую правую
	if begin < j {
		quicksort(arr, begin, j)
	}
	// и наоборот
	if end > i {
		quicksort(arr, i, end)
	}
}

func main() {
	arr := []int{20, 2, 6, 1, 0, 2, 5, -10, 2, 6, 1}
	beg := 0
	end := len(arr) - 1
	quicksort(arr, beg, end)
	fmt.Println(arr) // [-10 0 1 1 2 2 2 5 6 6 20]

}
