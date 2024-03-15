package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) GetDistance(ep Point) float64 {
	return math.Sqrt((ep.x-p.x)*(ep.x-p.x) + (ep.y-p.y)*(ep.y-p.y))
}

func main() {
	// объявляем две точки
	p1 := NewPoint(1, 3)
	p2 := NewPoint(2, 1)

	// получаем расстояние между ними
	fmt.Println(p1.GetDistance(p2))

}
