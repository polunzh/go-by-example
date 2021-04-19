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
	width, height float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) perim() float64 {
	return 3.14 * 2 * math.Pi
}

func measure(g geometry) {
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{r: 5}

	measure(r)
	measure(c)
}
