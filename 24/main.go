// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(nx, ny float64) *Point {
	return &Point{
		x: nx,
		y: ny,
	}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x - p.x, 2) + math.Pow(other.y - p.y, 2))
}

func main() {
	p1 := NewPoint(0, 1)
	p2 := NewPoint(2, -2)

	fmt.Println(p1.Distance(p2))
}