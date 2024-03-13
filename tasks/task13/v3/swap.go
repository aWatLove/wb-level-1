package main

import "fmt"

// Поменять местами два числа без создания временной переменной.
// 3-ий способ, с использованием сложения и вычитания
func main() {
	a := 1
	b := 5

	fmt.Println(a, b)

	a += b
	b = a - b
	a -= b

	fmt.Println(a, b)
}
