package main

import (
	"math"
	"math/cmplx"
)

/*
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
// represents a Unicode code point
float32 float64
complex64 complex128
 */

// Type conversions => The expression T(v) converts the value v to the type T.

import (
	"fmt"
)

// Constants cannot be declared using the := syntax.
// Constants can be character, string, boolean, or numeric values
const Pi = 3.14

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Type conversions => The expression T(v) converts the value v to the type T.
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Type inference
	// declaring a variable without specifying an explicit type GO internally specify its types based on value specified
	// i := 42           // int
	// f := 3.142        // float64
	// g := 0.867 + 0.5i // complex128
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("v is of type %T\n",string(v))
	fmt.Printf("v is of type %T\n", v)

	fmt.Println("Happy", Pi, "Day")
	//Pi = 3.1465 // error
}
