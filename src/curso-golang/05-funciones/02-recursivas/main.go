package main

import "fmt"

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	f := n * factorial(n-1)

	return f
}

func main() {
	// 3! = 3 * 2 * 1
	fmt.Println(factorial(3))
}
