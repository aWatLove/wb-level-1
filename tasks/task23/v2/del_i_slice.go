package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/
// 2-ой способ. Медленный, простой, не меняется порядок.
// Выполняется за линейное время, потому что происходит перезапись изза сдвига элементов
func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // слайс
	i := 4                                    // индекс элемента который нужно удалить
	fmt.Println(sl)                           // выводим слайс до удаления элемента
	sl = append(sl[:i], sl[i+1:]...)          // аппендим в подслайс до i-ого элемента, подслайс начинающийся с i+1 элемента
	fmt.Println(sl)                           // выводим слайс после удаления элемента
}
