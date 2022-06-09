package main

import (
	"fmt"
)

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

func linkedListOfIntegersExample() {
	l1 := IntList{1, nil}
  l2 := IntList{3, &l1}
	fmt.Println(l2.Sum())
}