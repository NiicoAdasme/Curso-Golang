package main

import "fmt"

func main() {
	var numeros [5] int

	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[4])

	// Array con valores
	nombres := [3] string {"Nicolas", "Carlos", "Alex"}

	fmt.Println(nombres)

	colores := [...] string{
		"rojo",
		"verde",
		"negro",
		"azul",
	}

	fmt.Println(colores, len(colores))

	// Indices definidos
	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}

	fmt.Println(monedas, len(monedas))

	// sub array
	// subMoneda := monedas[0:3]
	// subMoneda := monedas[:3]
	subMoneda := monedas[3:]
	fmt.Println(subMoneda)

}