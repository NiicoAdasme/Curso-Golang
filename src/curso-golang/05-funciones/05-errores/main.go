package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor int) (int, error){
	if divisor == 0{
		return 0, errors.New("No se puede dividir por cero")
	}else{
		return dividendo/divisor, nil
	}
}

func main() {
	result, error := division(4,2)

	if error == nil{
		fmt.Println(result)
	}else{
		fmt.Println(error)
	}

}