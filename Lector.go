package main

import ("fmt"
	"io/ioutil"
	"os")

func main() {
	
	var incluir string
	fmt.Scanln(&incluir)
	nuevotexto := "\n"+incluir
	fichero, err := os.OpenFile("Fichero.txt", os.O_APPEND,0777)
	muestraError(err)
	escribir,err2 := fichero.WriteString(nuevotexto)
	muestraError(err2)
	fichero.Close()
	fmt.Println(escribir)
	texto, errorFichero := ioutil.ReadFile("Fichero.txt")
	muestraError(errorFichero)
	fmt.Println(string(texto))
}
func muestraError(e error){
	if(e != nil){
		fmt.Println("No se encontro el fichero")
		panic(e)
	}
}