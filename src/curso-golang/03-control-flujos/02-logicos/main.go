package main

import "fmt"

func main() {
	
	// not
	fmt.Println(!true)  // res: false


	// and
	// fmt.Println(true && true) // true
	fmt.Println(true && false) // false
	
	// or
	fmt.Println(true || false) // true

	b := 2
	r := b == 2 || !(4 > b)
	// true

	fmt.Println(r)
}