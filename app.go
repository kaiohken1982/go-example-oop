package main

import (
	"fmt"
)

type Point struct{ X, Y float64 }

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	p.ScaleBy(2)

	fmt.Println(p)
	fmt.Println(p.X)
	fmt.Println(p.Y)
}