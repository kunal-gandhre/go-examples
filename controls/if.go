package main

import "fmt"

func main() {
	evenOrOddNumber(3)
	evenOrOddNumber(2)
}

func evenOrOddNumber(n int){
	if n % 2 == 0 {
		fmt.Printf("\n %d Number is Even", n)
	} else {
		fmt.Printf("\n %d Number is odd", n)
	}
}