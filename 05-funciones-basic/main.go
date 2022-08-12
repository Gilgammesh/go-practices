package main

import "fmt"

// Funcion con múltiples parámetros y retornando un valor
func sumar(a int, b int) int {
	return a + b
}

// Funcion con múltiples parámetros y retornando múltiples de valores
func evaluarSuma(a int, b int, c int) (int, bool) {
	sum := a + b
	var status bool
	if sum == c {
		status = true
	}
	return sum, status
}

func sumarInteractivo(numero ...int) (int, int) {
	sum := 0
	size := 0
	for index, num := range numero {
		sum = sum + num
		size = index + 1
	}
	return sum, size
}

func main() {
	num1 := 4
	num2 := 6
	fmt.Printf("La suma de los numeros %d + %d es igual a %d \n", num1, num2, sumar(num1, num2))

	sumtest := 12
	sum, status := evaluarSuma(num1, num2, sumtest)
	fmt.Printf("Es la suma de %d + %d igual a %d ? \n", num1, num2, sumtest)
	fmt.Printf("%t, la suma es igual a %d \n", status, sum)

	sumInt, size := sumarInteractivo(5, 4, 8, 3)
	fmt.Printf("La suma de los %d numeros es igual a %d \n", size, sumInt)

}
