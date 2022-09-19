package main

import "fmt"

func main() {
	// APP para restaurante
	// Dscto de 10% hasta 100 de consumo
	// Dscto de 20% hasta 200 de consumo
	// Dscto de 30% mayor a 200 de consumo
	// IGV de 19%
	var consumo, dscto float64
	var datosDescto string

	fmt.Println("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo <=100 {
		// descuento del 10%
		datosDescto = "10%"
		dscto = consumo * 0.10
	}else if consumo <= 200 {
		// descuento del 20%
		datosDescto = "20%"
		dscto = consumo * 0.20
	}else if consumo > 200{
		// descuento del 30%
		datosDescto = "30%"
		dscto = consumo * 0.30
	}else{
		fmt.Println("Error al ingresar el consumo")
	}

	// operaciones
	monto_dscto := consumo - dscto
	igv := monto_dscto * 0.19
	total_a_pagar := monto_dscto + igv

	fmt.Println("------ Factura de consumo --------")
	fmt.Println("Descuento que aplica: ", datosDescto)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Descuento: ", dscto)
	fmt.Println("IGV: ", igv)
	fmt.Println("Total a pagar: ", total_a_pagar)

}