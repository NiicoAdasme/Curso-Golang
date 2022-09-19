package main

import (
	"fmt"
)

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	// Agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)

	dias[7] ="SABADO"

	fmt.Println(dias)

	// Borrar dato
	delete(dias, 1)

	fmt.Println(dias, len(dias))

	// Nueva mapa
	estudiantes := make(map[string][]int)

	estudiantes["Alex"] = []int{13, 15, 16}
	estudiantes["Roel"] = []int{14, 13, 17}

	fmt.Println(estudiantes, len(estudiantes))

	fmt.Println(estudiantes["Alex"][1])

	




}