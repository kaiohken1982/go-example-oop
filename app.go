package main

import "fmt"

func VarDump(expression ...interface{} ) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

func main() {
	// linkedListOfIntegersExample()
	// pointExample()
  // valuesExample()
  // structEmbeddingExample()
	// cacheExample()
  intSetExample()
}