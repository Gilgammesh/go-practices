package main

import "fmt"

func main() {
	fmt.Println("Inicio")

	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(10, 6)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(3, 5)
	fmt.Println(result)
}

func sumar(a int, b int) int {
	return a + b
}

func restar(a int, b int) int {
	return a - b
}

func multiplicar(a int, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a int, b int) int {
		fmt.Println("Texto agregado")
		return f(a, b)
	}
}
