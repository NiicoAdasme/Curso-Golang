package main

import "fmt"

func saludar(nombre string) {
	// fmt.Println("Hola Mundo")
	fmt.Println("Hola,", nombre)

}

func sumar(a, b int) int {
	return a + b
}

func main() {
	// saludar("Nicolas")
	r := sumar(1,2)

	fmt.Println("La suma es:", r)
}