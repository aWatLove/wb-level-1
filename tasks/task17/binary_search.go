package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/
func binarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		midVal := arr[mid]

		if midVal == key {
			return mid
		} else if midVal < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{-5, 1, 4, 6, 8, 11, 15, 40, 62}
	key := 4
	fmt.Println("index:", binarySearch(arr, key))
}
