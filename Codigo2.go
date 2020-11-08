package main

import ("fmt"
		"os"
		"strconv"
		)

func main(){
	fmt.Println("****Mi programa con Go****")
	var edad int
	fmt.Println("Hola "+os.Args[1]+", Bienvenido al programa en GO")//se debe colar en el comando de ejecucion los valores a mostrar
	edad,_ = strconv.Atoi(os.Args[2]) //ejemplo "Go run Codigo2.go Pepe 18"
	if edad >= 18 && edad <=99 {
		fmt.Println("Eres mayor de edad")
	}else{
		fmt.Println("Eres menor de edad o demasiado mayor")}
}