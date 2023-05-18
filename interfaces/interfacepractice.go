package main

import (
	"fmt"
	"math"
)

type shape interface {
	Perimeter() float64
}

type Rectangle struct {
	length int
	width  int
}

func (r Rectangle) Perimeter() float64 {
	return float64(2 * (r.length + r.width))
}

type Triangle struct {
	length int
	width  int
	height int
}

type Circle struct {
	radious int
}

func (c Circle) vol() int {
	return c.radious * c.radious * c.radious
}

func printShape(s shape) {
	fmt.Printf("Type is %#v\n", s)
	fmt.Printf("perimeter is:: %v\n", s.Perimeter())
	fmt.Println()
}

func printVol(c Circle) {
	fmt.Printf("Volume is::%d", c.vol())
}

func (c Circle) Perimeter() float64 {
	return math.Pi * float64(2) * float64(c.radious)
}

func (t Triangle) Perimeter() float64 {
	return float64(t.width + t.height + t.length)
}

func main() {
	circe := Circle{
		radious: 5,
	}
	trian := Triangle{
		length: 5,
		width:  6,
		height: 7,
	}
	rect := Rectangle{
		length: 5,
		width:  6,
	}
	printShape(circe)
	printShape(rect)
	printShape(trian)
}
