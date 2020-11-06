package main

import (
	"fmt"
)

type persona struct{
	nombre string
	peso int
	altura float32
}

func main() {
	var yo = persona{
		"Alexis",
		65,
		1.67}
	
	fmt.Println(yo.nombre)
	var numeroX int = 10;
	var numeroY int = 6;

	//suma
	fmt.Println("la suma es:",numeroX+numeroY)

	//resta
	fmt.Println("la resta es:",numeroX-numeroY)

	//multiplicacion
	fmt.Println("la multiplicacion es:",numeroX*numeroY)

	//division
	fmt.Println("la division es:",numeroX/numeroY)

	//residuo
	fmt.Println("el residuo de la division es:",numeroX%numeroY)
}
