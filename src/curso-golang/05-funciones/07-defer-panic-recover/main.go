package main

import (
	"fmt"
	"os"
)

func main() {
	
	// Funcion
	defer func ()  {
		if error := recover(); error != nil{
			fmt.Println("Ups! El programa no finalizó de forma correcta")
		}
	}()

	if file, error := os.Open("holaa.txt"); error != nil{
		panic("No fue posible obtener el archivo")
	}else{
		
		defer func (){
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

}