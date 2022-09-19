package main

import "fmt"

func main() {
	numeros := []int{1,2,3}


	fmt.Println(numeros, len(numeros))

	// Agregar elementos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	// sub slicen
	subSlicen := numeros[:2]

	numeros[0] = 100

	fmt.Println("subSlicen: ", subSlicen, len(subSlicen))
	fmt.Println(numeros, len(numeros))

	// Caracteristicas de un Slicen:
	// Punteros
	// Longitud
	// Capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}