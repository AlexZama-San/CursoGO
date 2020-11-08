package main

import (
	"fmt"
	"time"
	)

func main() {
	/*var numer int 
	for i :=0; i<= 5; i++{
		fmt.Println("introduzca un numero")
		fmt.Scanln(&numer)
		if numer%2==0{
			fmt.Println(numer,"es par")
		}else{
			fmt.Println(numer,"es impar")
		}
	}*/
	//pruebaForEach()
	momento := time.Now()
	hoy := momento.Weekday()
	switch hoy {
		case 0:
			fmt.Println("hoy es domingo")
		case 1:
			fmt.Println("hoy es Lunes")
		case 2:
			fmt.Println("hoy es Martes")
		case 3:
			fmt.Println("hoy es Miercoles")
		case 4:
			fmt.Println("hoy es Jueves")
		case 5:
			fmt.Println("hoy es Viernes")
		case 6:
			fmt.Println("hoy es Sabado")
		default:
			fmt.Println("hoy es otro dia de la semana")
	}
}
func pruebaForEach(){
	coches:=[3]string{"Lamborguini","Bugatti","Ferrari"}
	for _,impcoche:= range coches{
		fmt.Printf("%s \n",impcoche)
	}
}