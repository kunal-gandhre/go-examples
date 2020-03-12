package main

import (
	"fmt"
	"math"
)

func main() {
	evenOrOddNumber(3)
	evenOrOddNumber(2)
	fmt.Println(pow(3, 2, 10))
}

func evenOrOddNumber(n int) {
	if n%2 == 0 {
		fmt.Printf("\n %d Number is Even", n)
	} else {
		fmt.Printf("\n %d Number is odd", n)
	}
}

// If with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
