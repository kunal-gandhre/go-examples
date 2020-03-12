package main

import "fmt"

// simple one
func add(x int, y int) int {
	return x + y
}

// short hand
func add1(x, y int) int {
	return x + y
}

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Simple => add : ", add(42, 13))
	fmt.Println("short hand => add2 : ", add1(42, 13))
	a, b := swap("GO", "Lang")
	fmt.Println("Multiple results => swap : ", a, b)
	x, y := split(45)
	fmt.Println("Named return values => split : ", x, y)
}
