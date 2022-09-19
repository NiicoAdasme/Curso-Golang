package main

import (
	"fmt"
)

// Struct persona
type Persona struct{
	nombre string
	edad int
}

// Metodos
func (p *Persona) imprimir(){
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editNombre(nuevoNombre string){
	p.nombre = nuevoNombre
}

// Herencia
type Empleado struct{
	Persona
	sueldo float64
}

func main() {
	p1 := Persona{"Nicolas", 24}

	// fmt.Println(p1)

	// p1.nombre = "Roel"

	// fmt.Println(p1)

	p2 := Persona{
		nombre: "Claudio",
		edad: 25,
	}

	// fmt.Println(p2)

	p1.imprimir()
	p2.imprimir()

	p1.editNombre("Roel")
	p1.imprimir()

	em1 := Empleado{
		sueldo: 500,
	}
	em1.nombre = "Pedro"
	em1.edad = 30
	em1.imprimir()
	fmt.Println(em1)


}