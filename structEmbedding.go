package main

import (
	"fmt"
  "image/color"
)

// type Point struct{ X, Y float64 }

type ColoredPoint struct {
  Point
  Color color.RGBA
}

func structEmbeddingExample() {

  var cp ColoredPoint
  cp.X = 1
  
  VarDump(cp)
  
  fmt.Println(cp.Point.X) // "1"
  cp.Point.Y = 2
  fmt.Println(cp.Y) // "2"


}