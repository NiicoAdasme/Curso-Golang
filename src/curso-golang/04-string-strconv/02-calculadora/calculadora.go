package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string )int{
	// Entrada: 4+4+5+8
	valores := strings.Split(expresion, "+")
	var result int

	for _, valor := range valores{
		num, error := strconv.Atoi(valor)
		if error != nil{
			fmt.Println(error)
			fmt.Println("Error: Debe ingresar numeros enteros y solo sumas")
		}else{
			result += num
		}
		
	}

	return result

}

func main() {
	var expresion string
	var result int

	fmt.Print("=>")
	fmt.Scanln(&expresion)

	result = sumar(expresion)

	fmt.Printf("=> %d", result)
}