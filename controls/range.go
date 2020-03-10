package main

import "fmt"

// Range is always used in conjunction with a data structure.

func main() {
	nums := []int{1,2,3,4,5,6}

	for _, num := range nums {
		fmt.Print(num)
	}

	str := "KUNAL"
	for k, v := range str {
		fmt.Printf("Key %d char %c : value %d \n", k, v, v)
	}
}