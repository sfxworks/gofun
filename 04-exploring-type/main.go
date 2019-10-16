package main

import (
	"fmt"
)

var y = 42

//Declare var of z is of type string
// assigning value "shaken not stirred"
var z = "Shaken, not stirred"

var a = `Yeeet said 
"Shaken, not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(z)
	fmt.Printf("%T\n", a)

	//z = 43  cannot assign different type like node/php/literally most all languages ive learned
	//Wow I kind of have an advantage here wee as3
	//fmt.Println(y)
	//fmt.Printf("%T\n", y)

}



//PRIMITIVE/built-in data types
//bool numeric string etc

//COMPOSITE/aggregate data types
//array struct slice etc