package main

import (
	"fmt"
)

type Point struct{ X, Y float64 }

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}


/**
 * In questo esempio il riferimento passato 
 * può essere nil, in questo caso è best practice 
 * specificare questa possibilità nell'implementazione ( if list == nil )
 */
type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list == nil { // qui: esplicitare che potrebbe arrivare nil
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	linkedListOfIntegersExample()
	// pointExample()
}

func linkedListOfIntegersExample() {
	l1 := IntList{1, nil}
  l2 := IntList{3, &l1}

  
	fmt.Println(l2.Sum())
}

func pointExample() {
	p := Point{1, 2}
	p.ScaleBy(2)

	fmt.Println(p)
	fmt.Println(p.X)
	fmt.Println(p.Y)
}