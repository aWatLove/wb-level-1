package main

import "fmt"

// Поменять местами два числа без создания временной переменной.
// 2-ой способ, с использованием XOR
func main() {
	a := 1
	b := 5

	fmt.Println(a, b)

	a ^= b
	b ^= a
	a ^= b

	fmt.Println(a, b)
}
