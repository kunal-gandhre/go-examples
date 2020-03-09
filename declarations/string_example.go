package main

import (
	"fmt"
	"strings"
)

// Declare constant
const Title string = "Title"
const Title2 = "Title 2"

// Declare package variable
var Country string = "INDIA"
// var Country = "INDIA" same as above

// Short assignment operator (:=) can’t be used for declare package level variables.
// CountryCode := "IN"

func main(){
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

	// ERROR : syntax error: non-declaration statement outside function body
	// fmt.Println("CountryCode:", CountryCode)

	str1, str2 := "Golang", "PROGRAM"
	// Convert to upper case
	fmt.Println("To Upper Case:", strings.ToUpper(str1))

	// Convert to lower case
	fmt.Println("To Lower Case:", strings.ToLower(str2))

	// String index
	//  index count from 0
	str3 := "Go"
	fmt.Printf("Index 0 value %c\n", str3[0])
}