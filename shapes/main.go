// interface hw
package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLangth float64
}

func main() {
	t := triangle{10, 10}
	s := square{10}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return (s.sideLangth * s.sideLangth)
}
