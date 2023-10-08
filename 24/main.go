package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func Distance(p1 *Point, p2 *Point) float64 {
	ySide := p1.y - p2.y
	xSide := p1.x - p2.x
	return math.Sqrt(math.Pow(ySide, 2) + math.Pow(xSide, 2))
}

func main() {
	p1 := New(10, 20)
	p2 := New(20, 10)
	fmt.Println(Distance(p1, p2))
}
