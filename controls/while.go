package main

// While loops are used when you are not sure how long code should be repeated.
// There is no do-while loop in Go

import "fmt"

func main() {
	i := 1
	max := 5
	// go does not have while keyword

	// technically go doesnt have while, but
	// for can be used while in go.
	for i < max {
		fmt.Println(i)
		i += 1
	}

	var input int = 2
	for ok := true; ok; ok = input != 2 {
		n, err := fmt.Scanln(&input)
		if n < 1 || err != nil {
			fmt.Println("invalid input")
			break
		}

		switch input {
		case 1:
			fmt.Println("hi")
		case 2:
			// Do nothing (we want to exit the loop)
			// In a real program this could be cleanup
		default:
			fmt.Println("def")
		}
	}
}
