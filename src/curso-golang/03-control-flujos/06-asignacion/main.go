package main

import "fmt"

func main() {
	a := 2
	// a = a + 2

	// Suma en asignacion
	a += 2
	fmt.Println("Suma: ",a)
	// Resta en asignacion
	a -= 2
	fmt.Println("Resta: ", a)
	// Multiplicacion en asignacion
	a *= 2
	fmt.Println("Multiplicacion: ",a)
	// Division en asignacion
	a /= 2
	fmt.Println("Division: ", a)
	// Modulo en asignacion
	a %= 2
	fmt.Println("Modulo: ", a)
}