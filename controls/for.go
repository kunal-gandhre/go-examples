package main

// For loop

import "fmt"

func main() {

	// using initialize variable
	fmt.Println("For loop 1")
	for x := 0; x < 4; x++ {
		fmt.Printf(" Iteration x: %d", x)
	}

	// using initialize variable
	fmt.Println("\n For loop 2")
	for x := 0; x < 4; {
		fmt.Printf(" Iteration x: %d", x)
		x++
	}

	// using inline variable
	sum := 1
	for sum < 50 {
		sum += sum
	}
	fmt.Println("\n value is ", sum)
}
