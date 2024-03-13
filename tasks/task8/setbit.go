package main

import (
	"errors"
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.
*/

func main() {
	var n int64 = 75  // число в котором хотим поменять бит
	var idx uint8 = 5 // индекс бита
	fmt.Printf("число:%d (%b)\nиндекс:%d\n", n, n, idx)
	n, err := setbit(n, idx) // вызываем функцию
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Новое число:%d (%b)", n, n)
}

// принимаемые значения: число, индекс бита. возвращаемые: число, ошибка
func setbit(num int64, bitIdx uint8) (int64, error) {
	if bitIdx < 0 || bitIdx > 63 { // проверяем что индекс в пределах [0-63]
		return 0, errors.New("The index is not within the range of 0-63")
	}

	var shift int64     // объявляем переменную для сдвига
	shift = 1 << bitIdx // совершаем побитовый сдвиг с шагом в bitIdx
	num = num ^ shift   // совершаем побитовую операцию XOR(исключающее или), чтобы изменить только один бит

	return num, nil
}
