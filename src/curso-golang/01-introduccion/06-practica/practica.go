package main

import "fmt"

func main() {
	var num1, num2 int
	
	fmt.Println("Ingrese un numero: ")
	fmt.Scanln(&num1)

	fmt.Println("Ingrese otro numero: ")
	fmt.Scanln(&num2)

	practica01(num1, num2)

	practica02(num1, num2)
}

func practica01 (n1, n2 int){
	result := n1 + n2
	fmt.Println("La suma es:", result)
}

func practica02 (n1, n2 int){
	result := n1 / n2
	fmt.Println("El cociente es:", result)
	
	result = n1 % n2
	fmt.Println("El residuo es:", result)
}