package main

import (
	"fmt"
)

//global variables

// Declare constant
const Title string = "Title"
const Title2 = "Title 2"

// Declare package variable
var Country string = "INDIA"
// var Country = "INDIA" same as above

// Short assignment operator (:=) can’t be used for declare package level variables.
// CountryCode := "IN"

func main(){

	// local variables below
	
	// Short assignment operator
	fname, lname := "Kunal", "Gandhre"
	age := 35

	// Print constant variable
	fmt.Println(Title)
	fmt.Println(Title2)

	// Print local variables
	fmt.Println("First Name:", fname)
	fmt.Println("Last Name:", lname)
	fmt.Println("Age:", age)

	// Print package variable
	fmt.Println("Country:", Country)

	// // variable names consist of letters, numbers, underscores, where the first letter cannot be numeric
	/*
	var (
		a int
		b bool
		s string
		apple int
		bee bool
		apple2 string
	)

	// When a variable is declared by a var, the system automatically gives it a zero value of that type:
	int i //= 0
	float f //= 0.0
	bool b //= False
	string s //= " "
	pointer p // = nil
	*/

	// simple declaration:
	a, b, c := 5, 7, "abc"
	fmt.Printf("%d %d %s\n",a,b, c)

	// nil value
	// var x = 1
	// cant add unused variable
	// so use blank identifier
	var _ = 1
	// x := nil error
	fmt.Println(nil)
}