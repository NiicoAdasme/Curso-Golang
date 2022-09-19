package main

import "fmt"

// Variables globales
var mensaje string

func function1(){
	mensaje = "Hola desde la funcion uno!"
	fmt.Println(mensaje)
}

func function2(){
	mensaje = "Hola desde la funcion dos!"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde la funcion principal"
	fmt.Println(mensaje)

	defer function1()
	function2()

	fmt.Println(mensaje)
}

