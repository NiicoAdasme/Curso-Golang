package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructuras
type Usuarios struct{
	UserName string
	Edad int
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request){
	// fmt.Fprintln(rw, "Hola mundo")
	template, err := template.ParseFiles("index.html")
	usuario := Usuarios{"Nicolas", 24}

	if err != nil{
		panic(err)
	}else{
		template.Execute(rw, usuario)
	}
}


func main() {
	// Archivos estaticos
	staticFile := http.FileServer(http.Dir("static"))


	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// Mux de static file
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))

	// Servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("Servidor esta funcionando en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}