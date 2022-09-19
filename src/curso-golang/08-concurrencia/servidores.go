package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, error := http.Get(servidor)
	if error != nil{
		// fmt.Println(servidor, "No esta disponible el servidor")
		canal <- servidor + "No esta disponible el servidor"
	}else{
		// fmt.Println(servidor, "Funcionando")
		canal <- servidor + "Funcionando"
	}
}

func main() {

	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"http://oregoom.com/",
		"http://www.udemy.com/",
		"http://www.youtube.com/",
		"http://www.facebook.com/",
		"http://www.google.com/",
	}

	for _, srv := range servidores {
		go revisarServidor(srv, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoQuePaso := time.Since(inicio)
	
	fmt.Println("Tiempo de ejecucion:", tiempoQuePaso)
}