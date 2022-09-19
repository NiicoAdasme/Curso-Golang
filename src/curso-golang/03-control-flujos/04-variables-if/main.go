package main

import "fmt"

func main() {

	if nombre, edad := "Roel", 26; nombre == "Alex" {
		fmt.Println("Hola ", nombre)
	}else{
		fmt.Printf("Nombre: %s - edad %d\n", nombre, edad)
	}


	// Obtener el valor de uin mapa
	users := make(map[string] string)

	users["Alex"] = "alex@gmail.com"
	users["Roel"] = "roel@gmail.com"

	fmt.Println(users["Alex"])

	// correo, ok := users["Juan"]
	// fmt.Println(correo, ok) 

	if correo, ok := users["Roel"]; ok {
		fmt.Println(correo)
	}else{
		fmt.Println("No fue posible obtener el valor")
	}
}