package main

import "fmt"

// Поменять местами два числа без создания временной переменной.
// 1-ый способ, с использованием синтаксической конструкции
func main() {
	a := 1
	b := 5

	fmt.Println(a, b)

	a, b = b, a

	fmt.Println(a, b)
}
