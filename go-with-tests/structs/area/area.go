package area

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	base   float64
	height float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}
