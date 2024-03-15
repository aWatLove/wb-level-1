package main

import (
	"errors"
	"fmt"
	"math/rand"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

	var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/

/*
в justString записывается результирующая строка, которая формируется из подсреза hugeString,
изза чего у нас не остается ссылки до внутреннего массива слайса hugeString и он не заполняет лишнюю память
*/

var justString string

func createHugeString(size int) []rune {
	baseString := []rune("пювбдыдарЛПДЫЛПТоРПАНЕК") // создаем слайс рун от базовой строки
	res := make([]rune, size)                       // результирующий срез

	for i := range res { // итеративно проходим по срезу и записываем рандомную руну в каждый индекс
		res[i] = baseString[rand.Intn(len(baseString))]
	}
	return res
}

func someFunc(minIdx, neededSize, maxSize int) error {
	if neededSize > maxSize {
		return errors.New("needed size cant be more than max size")
	}
	if maxSize <= 0 || neededSize <= 0 {
		return errors.New("size cant be negative or zero")
	}
	if minIdx < 0 {
		return errors.New("min index cant be negative")
	}

	hugeString := createHugeString(maxSize)
	justString = string(hugeString[minIdx:neededSize]) // присваиваем justString срез от мин до нужного значения
	return nil
}

func main() {
	if err := someFunc(50, 100, 150); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(justString) // дпыНбТывЛТдадПпПРПоКрАПдбддТКдбдпПдюРрдЫНТЛЛДаЫЛЛК
}
