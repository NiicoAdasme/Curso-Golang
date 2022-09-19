package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler
func Hola(rw http.ResponseWriter, r *http.Request){
	fmt.Println("El metodo es + " + r.Method)
	fmt.Fprintln(rw, "Hola mundo de Go web!!")
}

func pageNF(rw http.ResponseWriter, r *http.Request){
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request){
	http.Error(rw, "La pagina no funciona", http.StatusConflict)
}

func Saludar(rw http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())
	name := r.URL.Query().Get("name")
	rut := r.URL.Query().Get("rut")

	fmt.Fprintf(rw, "Hola %s tu rut es %s", name, rut)


}

func main() {

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", pageNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	// // Router
	// http.HandleFunc("/", Hola)
	// http.HandleFunc("/page", pageNF)
	// http.HandleFunc("/error", Error)
	// http.HandleFunc("/saludar", Saludar)

	// Crear servidor

	server := &http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}
	fmt.Println("Servidor esta funcionando en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	// log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())

}
