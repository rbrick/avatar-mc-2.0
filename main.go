package main

import "fmt"

func main() {

	one := Create(2)

	one.Set(0, 0, 1)
	one.Set(0, 1, 2)
	one.Set(1, 0, 3)
	one.Set(1, 1, 4)

	// scale
	one.Scale(2)

	fmt.Println(one.backingArray)
}
