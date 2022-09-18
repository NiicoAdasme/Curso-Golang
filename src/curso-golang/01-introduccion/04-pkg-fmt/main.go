package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)

	fmt.Println(mundo)

	nombre := "Alex"
	edad := 24

	fmt.Printf("Hola, %s y su edad es %d\n", nombre, edad)

	// %v para imprimir variables que no se sabe su tipo de dato
	fmt.Printf("Hola, %v y su edad es %d\n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %s y su edad es %d", nombre, edad)

	fmt.Println(mensaje)

	// %T devuelve el tipo de dato de la variable
	fmt.Printf("nombre: %T \n", edad)

	fmt.Print("Ingrese un nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Su nombre es: ", nombre)

}