package strinter

import "math"

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radious float64
}

type shape interface {
	Area() float64
}

func Perimeter(rt Rectangle) float64 {
	return 2 * (rt.height + rt.width)
}

func (rt Rectangle) Area() float64 {
	return rt.height * rt.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radious * c.radious
}
