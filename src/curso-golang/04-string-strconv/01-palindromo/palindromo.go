package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string)string{
	
	arrayCadena := strings.Split(cadena, "")
	arrayReverse := make([]string, 0)

	for i := len(arrayCadena) -1; i >= 0; i-- {
		arrayReverse = append(arrayReverse, arrayCadena[i])
	}

	return strings.Join(arrayReverse, "")
}


func esPalindromo(palabra string)bool{
	// Oso
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")
	
	palabraInvertida := reverse(palabra)

	return palabra == palabraInvertida
}

func main() {

	// var palabra string

	// fmt.Println("Ingrese una palabra: ")
	// fmt.Scanln(&palabra)
	
	if esPalindromo("Ana") {
		fmt.Println("Es palindromo")
	}else{
		fmt.Println("No es palindromo")
	}
	// reverse("Luz azul")
}