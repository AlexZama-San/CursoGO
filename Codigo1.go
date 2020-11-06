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
/*	var yo = persona{
		"Alexis",
		65,
		1.67}
	
	fmt.Println(yo.nombre)*/
	/*var numeroX float32 = 10;
	var numeroY float32 = 6;
	fmt.Println("Calculadora 1")
	Calculadora(numeroX,numeroY)*/
	fmt.Println(personas(20,21))

	lenguajes("Español","Ingles","Japones")

}


	func Calcula(nX float32, nY float32, op string) float32{
		var resultado float32
		if (op=="+"){
			resultado=nX+nY
		}else if(op=="-"){
			resultado=nX-nY
		}else if(op=="*"){
			resultado=nX*nY
		}else if(op=="/"){
			resultado=nX/nY
	}
	return resultado
	}
	func Calculadora(numeroX float32,numeroY float32){
			//suma
		fmt.Println("la suma es:",Calcula(numeroX,numeroY,"+"))

		//resta
		fmt.Println("la resta es:",Calcula(numeroX,numeroY,"-"))

		//multiplicacion
		fmt.Println("la multiplicacion es:",Calcula(numeroX,numeroY,"*"))

		//division
		fmt.Println("la division es:",Calcula(numeroX,numeroY,"/"))
	}

	func personas(numpersonas int, edad int) (string){
		totalpersonas:=func() int{
			return numpersonas*10
		}
		var mensaje string=fmt.Sprintf("El numero de personas es %d con edad de %d",totalpersonas(),edad)
		return mensaje
	}

	func lenguajes(lenguajes ...string){
		for _,lenguaje := range lenguajes{
			fmt.Println(lenguaje)
		}
	}
