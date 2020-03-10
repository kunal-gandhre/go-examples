package main

import (
	"fmt"
)

// multiple returns values

func Swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x, y := "Pune", "India"
	fmt.Println("Before Swap:", x, y)

	x, y = Swap(x, y)
	fmt.Println("After Swap:", x, y)
}
