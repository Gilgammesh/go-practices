package main

import "fmt"

// Función anónima
var calculo func(int, int) int = func(num1 int, num2 int) int {
	// Inicialmente sumo los números
	return num1 + num2
}

func main() {
	num1 := 7
	num2 := 5
	fmt.Printf("La suma de %d + %d es igual a %d \n", num1, num2, calculo(num1, num2))

	// Redefinimos la función cálculo:
	calculo = func(num1 int, num2 int) int {
		// Ahora restamos los números
		return num1 - num2
	}
	fmt.Printf("La resta de %d - %d es igual a %d \n", num1, num2, calculo(num1, num2))

	// Redefinimos la función cálculo:
	calculo = func(num1 int, num2 int) int {
		// Ahora multiplicamos los números
		return num1 * num2
	}
	fmt.Printf("La multimplicación de %d x %d es igual a %d \n", num1, num2, calculo(num1, num2))

	operaciones()

	/* CLOSURES */
	numero := 2
	mitabla := tabla(numero)
	for i := 0; i < 10; i++ {
		fmt.Println(mitabla())
	}
}

func operaciones() {
	// Otra forma de declarar funciones anónimas
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 2
		return numero * secuencia
	}
}
