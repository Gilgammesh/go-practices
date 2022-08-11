package main

import (
	"fmt"
	"strconv"
)

var numero int   // Por defecto valor entero 0
var texto string // Por defecto una cadena vacia ""
var estado bool  // Por defecto false

func main() {
	var numero1 int = 3
	numero2 := 2
	var texto1 string = "texto1"
	texto2 := "texto2"
	var estado1 bool = false
	estado2 := true

	numero3, texto3, estado3 := 1, "texto3", false

	texto3 = strconv.Itoa(numero2)

	var pi float32 = 3.1416

	fmt.Println(numero, numero1, numero2)
	fmt.Println(texto, texto1, texto2)
	fmt.Println(estado, estado1, estado2)
	fmt.Println(numero3, texto3, estado3)

	fmt.Println(pi + float32(numero2))

}
