package main

import (
	"fmt"
	"math/big"
)

/*
 Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

// реализация с помощью math/big

// функция суммы
func sum(a, b *big.Int) *big.Int {
	var res big.Int
	res.Add(a, b)
	return &res
}

// функция вычитания
func sub(a, b *big.Int) *big.Int {
	var res big.Int
	res.Sub(a, b)
	return &res
}

// функция умножения
func mul(a, b *big.Int) *big.Int {
	var res big.Int
	res.Mul(a, b)
	return &res
}

// функция деления
func div(a, b *big.Int) *big.Int {
	var res big.Int
	res.Div(a, b)
	return &res
}

func main() {
	// объявляем две переменные типа big.Int
	a := new(big.Int)
	b := new(big.Int)
	// устанавливаем значения
	a.SetString("40000000000000000000000000000000000000000000000000000000000000000", 10)
	b.SetString("20000000000000000000000234567876534000654300450000000023456000", 10)

	// вывод
	fmt.Printf("sum=%s\n", sum(a, b).String())
	fmt.Printf("sub=%s\n", sub(a, b).String())
	fmt.Printf("mul=%s\n", mul(a, b).String())
	fmt.Printf("div=%s\n", div(a, b).String())
}
