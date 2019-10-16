package main

import (
	"fmt" //https://godoc.org/fmt
)

func main() {
	n, err := fmt.Println("Hello, playground", 42, true) //interfaces that accept anything. Reminds me of ... in as3
	// ...interface{}
	fmt.Println(n)
	fmt.Println(err)
	//can do _ in place of err to void a return
	//n, _ := fmt.Println("Hello, playground", 42, true)
}
