package main

import "fmt"

// & Ampersand se encarga de apunta a la dirección en memoria

// * Asterisco obtiene el valor de la referencia o puntero

func main() {
	color := "rojo"
	// Valor y referencia en memoria
	fmt.Println(color, &color)

	auto1 := "Toyota"
	auto2 := "Suzuki"
	fmt.Println(auto1, &auto1)
	fmt.Println(auto2, &auto2)
	auto3 := auto1
	fmt.Println(auto3, &auto3)

	auto1 = "Ford"
	fmt.Println(auto1, &auto1)
	fmt.Println(auto3, &auto3)

	auto4 := &auto2
	// Mostramos el valor del puntero
	fmt.Println(auto4, &auto4, *auto4)

	auto2 = "Honda"
	fmt.Println(auto4, &auto4, *auto4)
}
