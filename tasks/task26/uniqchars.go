package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func unique(str string) bool {
	rm := make(map[rune]struct{})   // объявляем мапу, чтобы проверять на уникальность руны из строки
	str = strings.ToLower(str)      // приводим строку к одному регистру
	for _, e := range []rune(str) { // проходимся в цикле по слайсу рун из строки
		_, ok := rm[e] // смотрим есть ли элемент в мапе
		if ok {        // ok == true, то в строке есть не уникальные символы
			return false // возвращаем falsr
		}
		rm[e] = struct{}{} // добавляем руну как ключ в мапу
	}

	return true // после цикла по всей строки возвращаем true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Println(s1, "-", unique(s1))
	fmt.Println(s2, "-", unique(s2))
	fmt.Println(s3, "-", unique(s3))
}
