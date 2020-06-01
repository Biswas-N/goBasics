package main

import (
	"fmt"
	"math"
)

// Interfaces can be an contract or min-requirements of
// methods existing to a struct, similar to interfaces in
// OOP languages
type shape interface {
	area() float64
}

type circle struct {
	x, y, radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

// Methods based on above interface, but are compatible
// with both circle and rectangle
func getArea(s shape) float64 {
	return s.area()
}

func main() {
	circle1 := circle{
		radius: 22,
		x:      0,
		y:      0,
	}
	rect1 := rectangle{
		height: 10,
		width:  30,
	}

	fmt.Printf("Circle area: %.2f\n", getArea(circle1))
	fmt.Println("Rectangle area: ", getArea(rect1))
}
