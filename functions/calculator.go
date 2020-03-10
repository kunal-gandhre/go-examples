package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func main() {
	x, y := 20, 10
	result := Add(x, y)
	fmt.Println("[Add]:", result)
	result = Subtract(x, y)
	PrintResult(result, Subtract )

	// Multiple return
	x, y = values()
	fmt.Println(x)
	fmt.Println(y)

	// Closure
	sequenceGenerator := makeSequence()
	fmt.Println(sequenceGenerator())
	fmt.Println(sequenceGenerator())

	// recursion
	fmt.Printf("Recursion factorial of 3 is %d", factorial(3))

	// function returning error
	_, err := checkError()
	fmt.Printf("\nError check %s", err)
	fmt.Println(checkError())
}

func PrintResult(result int, opp interface{}) {
	fmt.Printf("Calculation '%s': result is %d", runtime.FuncForPC(reflect.ValueOf(opp).Pointer()).Name(), result)
}

// Multiple return
func values() (int, int) {
	return 2,4
}

// Closure --> Nested functions
func makeSequence() func() int {
	i:=1
	return func() int {
		i+=2
		return i
	}
}

// Recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// function returning error
func checkError() (int, error) {
	return -1, errors.New("Something wrong")
}