// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) set(X, Y float64) {
	p.x, p.y = X, Y
}

func findDistance(p1, p2 Point) float64 {
	differenceX, differenceY := p1.x-p2.x, p1.y-p2.y
	if differenceX < 0 {
		differenceX = -differenceX
	}

	if differenceY < 0 {
		differenceY = -differenceY
	}

	return math.Sqrt(math.Pow(differenceX, 2) + math.Pow(differenceY, 2))
}

func main() {
	var p1, p2 Point
	p1.set(6, 8)
	p2.set(0, 0)
	fmt.Printf("First point:\t%v\nSecond point:\t%v\n", p1, p2)
	fmt.Printf("Distance besides two theese points: %v", findDistance(p1, p2))
}
