package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height int
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return float64(r.width) * float64(r.height)
}

func (r rect) perim() float64 {
	return 2*float64(r.width) + 2*float64(r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("cicrle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	detectCircle(r)
	detectCircle(c)

	measure(r)
	measure(c)
}
