package main

import "fmt"

func main() {

	var altura float32

	fmt.Println("Ingresa la altura en metros")
	fmt.Scanf("%f", &altura)

	fmt.Println(&altura)

	// convierto usando el valor
	fmt.Printf("La altura en pies es: %f \n", convertirMetrosEnPies(altura))

	// convierto usando el puntero
	convertirMetrosEnPiesPuntero(&altura)
	fmt.Printf("La altura en pies es: %f \n", altura)
}

// Convertir metros en pies usando el valor de altura
func convertirMetrosEnPies(altura float32) float32 {
	fmt.Printf("Primer método con valor: %f - %v \n", altura, &altura)
	altura = altura * 3.28
	return altura
}

// Convertir metros en pies usando el puntero de altura
func convertirMetrosEnPiesPuntero(altura *float32) {
	fmt.Printf("Segundo método con puntero: %v - %v, %f \n", altura, &altura, *altura)
	*altura = *altura * 3.28
}
