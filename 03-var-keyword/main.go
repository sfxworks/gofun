package main

import (
	"fmt"
)

//dec var y
//assign value
//dec & assign = initialization
var y = 43 //private or public var hum

//Declare var with identifier z
//set type of int
// assigns 0 to type int
var z int //0 value gets assigned to int

func main() {
	//short dec operator
	//Declare var and assign value at same time
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again", y)

}
