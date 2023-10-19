package main

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) DistanceTo(other Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(4, 5)

	distance := p1.DistanceTo(p2)
	fmt.Println("Расстояние между точками:", distance)
}
