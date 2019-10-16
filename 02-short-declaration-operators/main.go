//https://golang.org/ref/spec#Identifiers
//extended backus naur form
// identifier = letter {letter | unicode_digit } .

//:= Its a gofer ooo
//https://golang.org/ref/spec#Operators_and_punctuation

package main

import (
	"fmt"
)

func main() {
	x := 42 // Declare - new concept. Declare and assign.
	fmt.Println(x)
	x = 99 // Assignment
	fmt.Println(x)

	y := 100 + 24 //Declare + Assign + expression
	fmt.Println(y)
}
