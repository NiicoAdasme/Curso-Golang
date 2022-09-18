package main

import "fmt"

func main()  {
	a := 20
	b := 10

	// Suma
	result := a + b

	fmt.Println("Suma: ", result)

	// Resta
	result = a - b
	
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = a * b

	fmt.Println("Multi: ", result)

	// Division
	result = a / b

	fmt.Println("Divison: ", result)

	var div float64 = 3.0 / 2.0

	fmt.Println("Divison 3/2: ", div)

	// Modulo
	result = 3 % 2

	fmt.Println("Modulo: ", result)
}