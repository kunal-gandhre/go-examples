package main

// Create a function which returns the string changed to camel case.

import (
	"fmt"
	"strings"
	"unicode"
)

func main(){
	fmt.Printf("Input: firstName --> Output: %s \n", ToFirstUpper("firstName"))
	fmt.Printf("Input: first_name --> Output: %s \n", ToFirstUpper("first_name"))
	fmt.Printf("Input: user_first_name --> Output: %s \n", ToFirstUpper("user_first_name"))
}

// Returns the string changed to camel case
func ToFirstUpper(s string) string {
	var camelCase string = ""
	if !strings.Contains(s, "_") {
		return s
	}
	if len(s) < 1 { // if the empty string
		return s
	}
	// Trim the string
	t := strings.Trim(s, " ")
	// Convert all letters to lower case
	t = strings.ToLower(t)

	str := strings.Split(s, "_")
	for k, v := range str {
		if k == 0 {
			camelCase += v;
			continue
		}
		// character values from integer values
		res := []rune(v)
		// Convert first letter to upper case
		res[0] = unicode.ToUpper(res[0])
		camelCase += string(res);
	}
	return string(camelCase)
}
