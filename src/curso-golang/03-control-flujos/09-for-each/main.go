package main

import "fmt"

func main() {
	nombre := [...]string{"Alex", "Roel", "Juan"}

	for indice, elemento := range nombre {
		fmt.Println(indice, elemento)
	}
	
}