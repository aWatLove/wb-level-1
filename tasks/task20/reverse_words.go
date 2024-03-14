package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	str := "snow dog sun"
	words := strings.Split(str, " ") // разбиваем строку на срез строк с разделителем в виде пробела

	fmt.Println(str)            // выводим строку до переворачивания
	var builder strings.Builder // объявляем Builder

	// проходимся по срезу слов с конца и добавляем значение в билдер
	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i])
		if i < len(words)+1 { // если не последний элемент, добавляем пробел
			builder.WriteString(" ")
		}
	}
	str = builder.String() // формируем строк из builder'а
	fmt.Println(str)
}
