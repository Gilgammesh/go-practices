package main

import "fmt"

var tabla [10]int
var matriz [3][3]int

func main() {
	for i := 0; i < len(tabla); i++ {
		tabla[i] = i + 1
	}
	fmt.Println(tabla)

	tabla1 := [6]string{"6", "5", "4", "3", "2", "1"}
	for _, item := range tabla1 {
		fmt.Println(item)
	}

	matriz[2][0] = 1
	matriz[2][1] = 4
	matriz[2][2] = 2
	fmt.Println(matriz)

	/* SLICES (Matrices o vectores dinámicos) */
	slice1 := []int{2, 5, 6}
	fmt.Println(slice1)

	variante2()

	variante3()

	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	slice2 := elementos[3:4]
	fmt.Println(slice2)
}

func variante3() {
	// Función que reserva n elementos en memoria
	slice3 := make([]int, 5, 20)
	fmt.Printf("Longitud: %d - Capacidad: %d \n", len(slice3), cap(slice3))
}

func variante4() {
	slice4 := make([]int, 0, 0)
	for i := 0; i < 10; i++ {
		slice4 = append(slice4, i)
	}
	fmt.Println(slice4)
	fmt.Printf("Longitud: %d - Capacidad: %d \n", len(slice4), cap(slice4))
}
