package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rect struct {
	width, height float64
}

func (c Circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func (r Rect) area() float64 {
	return 2 * r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{2, 3, 3.5}

	rect := Rect{2, 3}

	fmt.Println(circle.area())
	fmt.Println(rect.area())

	fmt.Println("area of circle", getArea(circle))
	fmt.Println("area of rect", getArea(rect))

	fmt.Println("hello world")
}
