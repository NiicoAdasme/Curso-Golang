package main

import (
	"fmt"
	"math"
)

type Figura interface {
	area() float64
	perimetro() float64
}

type Cuadrado struct {
	lado float64
}

type Circulo struct {
	radio float64
}

// Cuadrado
func (cua *Cuadrado) area() float64 {
	return cua.lado * cua.lado
}

func (cua *Cuadrado) perimetro() float64 {
	return 4*cua.lado
}

// Circulo
func (circ *Circulo) area() float64 {
	return math.Pi * (circ.radio * circ.radio)
}

func (circ *Circulo) perimetro() float64 {
	return 2 * math.Pi * circ.radio
}

func medidas(figura Figura){
	// fmt.Println(figura)
	fmt.Println("Area:", figura.area())
	fmt.Println("Perimetro:", figura.perimetro())
}

func main() {
	
	cuadrado := Cuadrado{lado: 4}
	fmt.Println("Cuadrado")
	medidas(&cuadrado)


	circulo := Circulo{radio: 5}
	fmt.Println("Circulo")
	medidas(&circulo)

}